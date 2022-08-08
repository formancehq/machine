package vm

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"testing"

	ledger "github.com/numary/ledger/pkg/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm/program"
	"github.com/stretchr/testify/assert"
)

const (
	DEBUG bool = false
)

type CaseResult struct {
	Printed  []core.Value
	Postings []ledger.Posting
	Metadata map[string]core.Value
	ExitCode byte
	Error    string
}

type TestCase struct {
	program  *program.Program
	vars     map[string]core.Value
	meta     map[string]map[string]core.Value
	balances map[string]map[string]int64
	expected CaseResult
}

func NewTestCase() TestCase {
	return TestCase{
		vars:     make(map[string]core.Value),
		meta:     make(map[string]map[string]core.Value),
		balances: make(map[string]map[string]int64),
		expected: CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			Metadata: make(map[string]core.Value),
			ExitCode: EXIT_OK,
			Error:    "",
		},
	}
}

func (c *TestCase) compile(t *testing.T, code string) {
	p, err := compiler.Compile(code)
	if err != nil {
		t.Fatalf("compile error: %v", err)
		return
	}
	c.program = p
}

func (c *TestCase) setVarsFromJSON(t *testing.T, str string) {
	var jvars map[string]json.RawMessage
	err := json.Unmarshal([]byte(str), &jvars)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	v, err := c.program.ParseVariablesJSON(jvars)
	if err != nil {
		t.Fatalf("parse variables error: %v", err)
	}
	c.vars = v
}

func (c *TestCase) setBalance(t *testing.T, account, asset string, amount int64) {
	if _, ok := c.balances[account]; !ok {
		c.balances[account] = make(map[string]int64)
	}
	c.balances[account][asset] = amount
}

func test(
	t *testing.T,
	test_case TestCase,
) {
	testimpl(t, test_case.program, test_case.expected, func(m *Machine) (byte, error) {
		err := m.SetVars(test_case.vars)
		if err != nil {
			return 0, err
		}
		{
			ch, err := m.ResolveResources()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				val := test_case.meta[req.Account][req.Key]
				if req.Error != nil {
					panic(req.Error)
				}
				req.Response <- val
			}
		}
		{
			ch, err := m.ResolveBalances()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				val := test_case.balances[req.Account][req.Asset]
				if req.Error != nil {
					panic(req.Error)
				}
				req.Response <- val
			}
		}
		return m.Execute()
	})
}

func testimpl(t *testing.T, prog *program.Program, expected CaseResult, exec func(*Machine) (byte, error)) {
	printed := []core.Value{}

	var wg sync.WaitGroup
	wg.Add(1)

	if prog == nil {
		t.Fatal("no program")
	}

	machine := NewMachine(*prog)
	machine.Debug = DEBUG
	machine.Printer = func(c chan core.Value) {
		for v := range c {
			printed = append(printed, v)
		}
		wg.Done()
	}
	exit_code, err := exec(machine)

	if err != nil && expected.Error != "" {
		if !strings.Contains(err.Error(), expected.Error) {
			t.Error(fmt.Errorf("unexpected execution error: %v", err))
			return
		} else {
			return
		}
	} else if err != nil {
		t.Error(fmt.Errorf("did not expect an execution error: %v", err))
		return
	} else if expected.Error != "" {
		t.Error(fmt.Errorf("expected an execution error"))
		return
	}

	if exit_code != expected.ExitCode {
		t.Error(fmt.Errorf("unexpected exit code: %v", exit_code))
		return
	}
	if exit_code != EXIT_OK {
		return
	}

	if expected.Postings == nil {
		expected.Postings = make([]ledger.Posting, 0)
	}
	if expected.Metadata == nil {
		expected.Metadata = make(map[string]core.Value)
	}

	assert.Equalf(t, expected.Postings, machine.Postings, "unexpected postings output: %v", machine.Postings)
	assert.Equalf(t, expected.Metadata, machine.TxMeta, "unexpected metadata output: %v", machine.TxMeta)

	wg.Wait()

	assert.Equalf(t, expected.Printed, printed, "unexpected metadata output: %v", printed)
}

func TestFail(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, "fail")
	tc.expected = CaseResult{
		Printed:  []core.Value{},
		Postings: []ledger.Posting{},
		ExitCode: EXIT_FAIL,
	}
	test(t, tc)
}

func TestPrint(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, "print 29 + 15 - 2")
	tc.expected = CaseResult{
		Printed:  []core.Value{core.Number(*big.NewInt(42))},
		Postings: []ledger.Posting{},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSend(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [EUR/2 100] (
		source=@alice
		destination=@bob
	)`)
	tc.setBalance(t, "alice", "EUR/2", 100)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "EUR/2",
				Amount:      ledger.NewMonetaryInt(100),
				Source:      "alice",
				Destination: "bob",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestVariables(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $rider
		account $driver
	}
	send [EUR/2 999] (
		source=$rider
		destination=$driver
	)`)
	tc.vars = map[string]core.Value{
		"rider":  core.Account("users:001"),
		"driver": core.Account("users:002"),
	}
	tc.setBalance(t, "users:001", "EUR/2", 1000)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "EUR/2",
				Amount:      ledger.NewMonetaryInt(999),
				Source:      "users:001",
				Destination: "users:002",
			},
		},
		ExitCode: EXIT_OK,
	}
}

func TestVariablesJSON(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $rider
		account $driver
		string 	$description
	}
	send [EUR/2 999] (
		source=$rider
		destination=$driver
	)
	set_tx_meta("description", $description)`)
	tc.setVarsFromJSON(t, `{
		"rider": "users:001",
		"driver": "users:002",
		"description": "midnight ride"
	}`)
	tc.setBalance(t, "users:001", "EUR/2", 1000)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "EUR/2",
				Amount:      ledger.NewMonetaryInt(999),
				Source:      "users:001",
				Destination: "users:002",
			},
		},
		Metadata: map[string]core.Value{
			"description": core.String("midnight ride"),
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSource(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $balance
		account $payment
		account $seller
	}
	send [GEM 15] (
		source = {
			$balance
			$payment
		}
		destination = $seller
	)`)
	tc.setVarsFromJSON(t, `{
		"balance": "users:001",
		"payment": "payments:001",
		"seller": "users:002"
	}`)
	tc.setBalance(t, "users:001", "GEM", 3)
	tc.setBalance(t, "payments:001", "GEM", 12)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(3),
				Source:      "users:001",
				Destination: "users:002",
			},
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(12),
				Source:      "payments:001",
				Destination: "users:002",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestAllocation(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $rider
		account $driver
	}
	send [GEM 15] (
		source = $rider
		destination = {
			80% to $driver
			8% to @a
			12% to @b
		}
	)`)
	tc.setVarsFromJSON(t, `{
		"rider": "users:001",
		"driver": "users:002"
	}`)
	tc.setBalance(t, "users:001", "GEM", 15)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(13),
				Source:      "users:001",
				Destination: "users:002",
			},
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(1),
				Source:      "users:001",
				Destination: "a",
			},
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(1),
				Source:      "users:001",
				Destination: "b",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestDynamicAllocation(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		portion $p
	}
	send [GEM 15] (
		source = @a
		destination = {
			80% to @b
			$p to @c
			remaining to @d
		}
	)`)
	tc.setVarsFromJSON(t, `{
		"p": "15%"
	}`)
	tc.setBalance(t, "a", "GEM", 15)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(13),
				Source:      "a",
				Destination: "b",
			},
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(2),
				Source:      "a",
				Destination: "c",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSendAll(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [USD/2 *] (
		source = @users:001
		destination = @platform
	)`)
	tc.setBalance(t, "users:001", "USD/2", 17)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "USD/2",
				Amount:      ledger.NewMonetaryInt(17),
				Source:      "users:001",
				Destination: "platform",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSendAllMulti(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [USD/2 *] (
		source = {
		  @users:001:wallet
		  @users:001:credit
		}
		destination = @platform
	)
	`)
	tc.setBalance(t, "users:001:wallet", "USD/2", 19)
	tc.setBalance(t, "users:001:credit", "USD/2", 22)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "USD/2",
				Amount:      ledger.NewMonetaryInt(19),
				Source:      "users:001:wallet",
				Destination: "platform",
			},
			{
				Asset:       "USD/2",
				Amount:      ledger.NewMonetaryInt(22),
				Source:      "users:001:credit",
				Destination: "platform",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestInsufficientFunds(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $balance
		account $payment
		account $seller
	}
	send [GEM 16] (
		source = {
			$balance
			$payment
		}
		destination = $seller
	)`)
	tc.setVarsFromJSON(t, `{
		"balance": "users:001",
		"payment": "payments:001",
		"seller": "users:002"
	}`)
	tc.setBalance(t, "users:001", "GEM", 3)
	tc.setBalance(t, "payments:001", "GEM", 12)
	tc.expected = CaseResult{
		Printed:  []core.Value{},
		Postings: []ledger.Posting{},
		ExitCode: EXIT_FAIL_INSUFFICIENT_FUNDS,
	}
	test(t, tc)
}

func TestWorldSource(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [GEM 15] (
		source = {
			@a
			@world
		}
		destination = @b
	)`)
	tc.setBalance(t, "a", "GEM", 1)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(1),
				Source:      "a",
				Destination: "b",
			},
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(14),
				Source:      "world",
				Destination: "b",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestNoEmptyPostings(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [GEM 2] (
		source = @world
		destination = {
			90% to @a
			10% to @b
		}
	)`)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "GEM",
				Amount:      ledger.NewMonetaryInt(2),
				Source:      "world",
				Destination: "a",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestNoEmptyPostings2(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [GEM *] (
		source = @foo
		destination = @bar
	)`)
	tc.setBalance(t, "foo", "GEM", 0)
	tc.expected = CaseResult{
		Printed:  []core.Value{},
		Postings: []ledger.Posting{},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestAllocateDontTakeTooMuch(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [CREDIT 200] (
		source = {
			@users:001
			@users:002
		}
		destination = {
			1/2 to @foo
			1/2 to @bar
		}
	)`)
	tc.setBalance(t, "users:001", "CREDIT", 100)
	tc.setBalance(t, "users:002", "CREDIT", 110)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "CREDIT",
				Amount:      ledger.NewMonetaryInt(100),
				Source:      "users:001",
				Destination: "foo",
			},
			{
				Asset:       "CREDIT",
				Amount:      ledger.NewMonetaryInt(100),
				Source:      "users:002",
				Destination: "bar",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestMetadata(t *testing.T) {
	commission, _ := core.NewPortionSpecific(*big.NewRat(125, 1000))
	tc := NewTestCase()
	tc.compile(t, `vars {
		account $sale
		account $seller = meta($sale, "seller")
		portion $commission = meta($seller, "commission")
	}
	send [EUR/2 100] (
		source = $sale
		destination = {
			remaining to $seller
			$commission to @platform
		}
	)`)
	tc.setVarsFromJSON(t, `{
		"sale": "sales:042"
	}`)
	tc.meta = map[string]map[string]core.Value{
		"sales:042": {
			"seller": core.Account("users:053"),
		},
		"users:053": {
			"commission": *commission,
		},
	}
	tc.setBalance(t, "sales:042", "EUR/2", 2500)
	tc.setBalance(t, "users:053", "EUR/2", 500)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "EUR/2",
				Amount:      ledger.NewMonetaryInt(88),
				Source:      "sales:042",
				Destination: "users:053",
			},
			{
				Asset:       "EUR/2",
				Amount:      ledger.NewMonetaryInt(12),
				Source:      "sales:042",
				Destination: "platform",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestTrackBalances(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `
	send [COIN 50] (
		source = @world
		destination = @a
	)
	send [COIN 100] (
		source = @a
		destination = @b
	)`)
	tc.setBalance(t, "a", "COIN", 50)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(50),
				Source:      "world",
				Destination: "a",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(100),
				Source:      "a",
				Destination: "b",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestTrackBalances2(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `
	send [COIN 50] (
		source = @a
		destination = @z
	)
	send [COIN 50] (
		source = @a
		destination = @z
	)`)
	tc.setBalance(t, "a", "COIN", 60)
	tc.expected = CaseResult{
		Printed:  []core.Value{},
		Postings: []ledger.Posting{},
		ExitCode: EXIT_FAIL_INSUFFICIENT_FUNDS,
	}
	test(t, tc)
}

func TestTrackBalances3(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [COIN *] (
		source = @foo
		destination = {
			max [COIN 1000] to @bar
			remaining kept
		}
	)
	send [COIN *] (
		source = @foo
		destination = @bar
	)`)
	tc.setBalance(t, "foo", "COIN", 2000)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(1000),
				Source:      "foo",
				Destination: "bar",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(1000),
				Source:      "foo",
				Destination: "bar",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSourceAllotment(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [COIN 100] (
		source = {
			60% from @a
			35.5% from @b
			4.5% from @c
		}
		destination = @d
	)`)
	tc.setBalance(t, "a", "COIN", 100)
	tc.setBalance(t, "b", "COIN", 100)
	tc.setBalance(t, "c", "COIN", 100)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(61),
				Source:      "a",
				Destination: "d",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(35),
				Source:      "b",
				Destination: "d",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(4),
				Source:      "c",
				Destination: "d",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSourceOverlapping(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [COIN 99] (
		source = {
			15% from {
				@b
				@a
			}
			30% from @a
			remaining from @a
		}
		destination = @world
	)`)
	tc.setBalance(t, "a", "COIN", 99)
	tc.setBalance(t, "b", "COIN", 3)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(3),
				Source:      "b",
				Destination: "world",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(96),
				Source:      "a",
				Destination: "world",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestSourceComplex(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `vars {
		monetary $max
	}
	send [COIN 200] (
		source = {
			50% from {
				max [COIN 4] from @a
				@b
				@c
			}
			remaining from max $max from @d
		}
		destination = @platform
	)`)
	tc.setVarsFromJSON(t, `{
		"max": {
			"asset": "COIN",
			"amount": 120
		}
	}`)
	tc.setBalance(t, "a", "COIN", 1000)
	tc.setBalance(t, "b", "COIN", 40)
	tc.setBalance(t, "c", "COIN", 1000)
	tc.setBalance(t, "d", "COIN", 1000)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(4),
				Source:      "a",
				Destination: "platform",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(40),
				Source:      "b",
				Destination: "platform",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(56),
				Source:      "c",
				Destination: "platform",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(100),
				Source:      "d",
				Destination: "platform",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestDestinationComplex(t *testing.T) {
	tc := NewTestCase()
	tc.compile(t, `send [COIN 100] (
		source = @world
		destination = {
			20% to @a
			20% kept
			60% to {
				max [COIN 10] to @b
				remaining to @c
			}
		}
	)`)
	tc.expected = CaseResult{
		Printed: []core.Value{},
		Postings: []ledger.Posting{
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(20),
				Source:      "world",
				Destination: "a",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(10),
				Source:      "world",
				Destination: "b",
			},
			{
				Asset:       "COIN",
				Amount:      ledger.NewMonetaryInt(50),
				Source:      "world",
				Destination: "c",
			},
		},
		ExitCode: EXIT_OK,
	}
	test(t, tc)
}

func TestNeededBalances(t *testing.T) {
	p, err := compiler.Compile(`vars {
		account $a
	}
	send [GEM 15] (
		source = {
			$a
			@b
			@world
		}
		destination = @c
	)`)

	if err != nil {
		t.Fatalf("did not expect error on Compile, got: %v", err)
	}

	m := NewMachine(*p)

	err = m.SetVars(map[string]core.Value{
		"a": core.Account("a"),
	})
	if err != nil {
		t.Fatalf("did not expect error on SetVars, got: %v", err)
	}
	{
		ch, err := m.ResolveResources()
		if err != nil {
			t.Fatalf("did not expect error on ResolveResources, got: %v", err)
		}
		for range ch {
			t.Fatalf("did not expect to need any metadata")
		}
	}

	expected := map[string]map[string]struct{}{
		"a": {
			"GEM": {},
		},
		"b": {
			"GEM": {},
		},
	}
	{
		ch, err := m.ResolveBalances()
		if err != nil {
			t.Fatalf("did not expect error on ResolveBalances, got: %v", err)
		}
		for req := range ch {
			if req.Error != nil {
				t.Fatalf("did not expect error in balance request: %v", req.Error)
			}
			if _, ok := expected[req.Account][req.Asset]; ok {
				delete(expected[req.Account], req.Asset)
				if len(expected[req.Account]) == 0 {
					delete(expected, req.Account)
				}
				req.Response <- 0
			} else {
				t.Fatalf("did not expect to need %v balance of %v", req.Asset, req.Account)
			}
		}
	}
	if len(expected) != 0 {
		t.Fatalf("some balances were not requested: %v", expected)
	}
}

func TestSetTxMeta(t *testing.T) {
	p, err := compiler.Compile(`
	set_tx_meta("aaa", @platform)
	set_tx_meta("bbb", GEM)
	set_tx_meta("ccc", 45)
	set_tx_meta("ddd", "hello")
	set_tx_meta("eee", [COIN 30])
	`)

	if err != nil {
		t.Fatalf("did not expect error on Compile, got: %v", err)
	}

	m := NewMachine(*p)

	{
		ch, _ := m.ResolveResources()
		for range ch {
		}
	}

	{
		ch, _ := m.ResolveBalances()
		for range ch {
		}
	}

	_, err = m.Execute()

	if err != nil {
		t.Fatalf("did not expect error on Execute, got: %v", err)
	}

	expected_meta := map[string]json.RawMessage{
		"aaa": json.RawMessage(`{"type":"account","value":"platform"}`),
		"bbb": json.RawMessage(`{"type":"asset","value":"GEM"}`),
		"ccc": json.RawMessage(`{"type":"number","value":45}`),
		"ddd": json.RawMessage(`{"type":"string","value":"hello"}`),
		"eee": json.RawMessage(`{"type":"monetary","value":{"asset":"COIN","amount":30}}`),
	}

	meta := m.GetTxMetaJson()

	fmt.Printf("%v", len(meta))

	if len(meta) != 5 {
		t.Fatalf("unexpected transaction metadata")
	}

	for k, v := range meta {
		if expected, ok := expected_meta[k]; ok {
			if rm, ok := v.(json.RawMessage); ok {
				if string(rm) != string(expected) {
					fmt.Printf("%v\n", string(rm))
					fmt.Printf("%v\n", string(expected))
					t.Fatalf("unexpected transaction metadata")
				}
			}
		}
	}
}
