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
	Code      string
	Variables map[string]core.Value
	Expected  CaseResult
}

type TestCaseJSON struct {
	Code      string
	Variables map[string]core.Value
	Expected  CaseResult
}

func test(
	t *testing.T,
	code string,
	variables map[string]core.Value,
	meta map[string]map[string]core.Value,
	balances map[string]map[string]uint64,
	expected CaseResult,
) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		err := m.SetVars(variables)
		if err != nil {
			return 0, err
		}
		{
			ch, err := m.ResolveResources()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				val := meta[req.Account][req.Key]
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
				val := balances[req.Account][req.Asset]
				if req.Error != nil {
					panic(req.Error)
				}
				req.Response <- val
			}
		}
		return m.Execute()
	})
}

func testJSON(
	t *testing.T,
	code string,
	variables string,
	meta map[string]map[string]core.Value,
	balances map[string]map[string]uint64,
	expected CaseResult,
) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		var v map[string]json.RawMessage
		err := json.Unmarshal([]byte(variables), &v)
		if err != nil {
			return 0, err
		}
		err = m.SetVarsFromJSON(v)
		if err != nil {
			return 0, err
		}
		{
			ch, err := m.ResolveResources()
			if err != nil {
				return 0, err
			}
			for req := range ch {
				if req.Error != nil {
					panic(req.Error)
				}
				val, ok := meta[req.Account][req.Key]
				if !ok {
					t.Fatalf("case error: missing metadata %q of %v", req.Key, req.Account)
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
				if req.Error != nil {
					panic(req.Error)
				}
				val, ok := balances[req.Account][req.Asset]
				if !ok {
					t.Fatalf("case error: missing %v balance of %v", req.Asset, req.Account)
				}
				req.Response <- val
			}
		}
		return m.Execute()
	})
}

func testimpl(t *testing.T, code string, expected CaseResult, exec func(*Machine) (byte, error)) {
	p, err := compiler.Compile(code)

	if err != nil {
		t.Error(fmt.Errorf("compile error: %v", err))
		return
	}

	printed := []core.Value{}

	var wg sync.WaitGroup
	wg.Add(1)

	machine := NewMachine(p)
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

	if len(machine.Postings) != len(expected.Postings) {
		t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings))
		return
	} else {
		for i := range machine.Postings {
			if machine.Postings[i] != expected.Postings[i] {
				t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings[i]))
				return
			}
		}
	}

	if len(machine.TxMeta) != len(expected.Metadata) {
		t.Error(fmt.Errorf("unexpected metadata output: %v", machine.TxMeta))
		return
	} else {
		for k := range machine.TxMeta {
			if machine.TxMeta[k] != expected.Metadata[k] {
				t.Error(fmt.Errorf(
					"unexpected metadata output value for key %s, got: %v, expected: %v",
					k,
					machine.TxMeta[k],
					expected.Metadata[k],
				))
				return
			}
		}
	}

	wg.Wait()

	if len(printed) != len(expected.Printed) {
		t.Error(fmt.Errorf("unexpected print output: %v", printed))
		return
	} else {
		for i := range printed {
			if printed[i] != expected.Printed[i] {
				t.Error(fmt.Errorf("unexpected print output: %v", printed[i]))
				return
			}
		}
	}
}

func TestFail(t *testing.T) {
	test(t,
		"fail",
		map[string]core.Value{},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_FAIL,
		},
	)
}

func TestPrint(t *testing.T) {
	test(t,
		"print 29 + 15 - 2",
		map[string]core.Value{},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{
			Printed:  []core.Value{core.Number(42)},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSend(t *testing.T) {
	test(t,
		`send [EUR/2 100] (
			source=@alice
			destination=@bob
		)`,
		map[string]core.Value{},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"alice": {
				"EUR/2": 100,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      100,
					Source:      "alice",
					Destination: "bob",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestVariables(t *testing.T) {
	test(t,
		`vars {
			account $rider
			account $driver
		}
		send [EUR/2 999] (
			source=$rider
			destination=$driver
		)`,
		map[string]core.Value{
			"rider":  core.Account("users:001"),
			"driver": core.Account("users:002"),
		},
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"EUR/2": 1000,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      999,
					Source:      "users:001",
					Destination: "users:002",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestVariablesJSON(t *testing.T) {
	testJSON(t,
		`vars {
			account $rider
			account $driver
			string 	$description
		}
		send [EUR/2 999] (
			source=$rider
			destination=$driver
		)
		set_tx_meta("description", $description)`,
		`{
			"rider": "users:001",
			"driver": "users:002",
			"description": "midnight ride"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"EUR/2": 1000,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      999,
					Source:      "users:001",
					Destination: "users:002",
				},
			},
			Metadata: map[string]core.Value{
				"description": core.String("midnight ride"),
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestVariablesJSONInvalid(t *testing.T) {
	testJSON(t,
		`vars {
			portion $p
		}
		send [EUR/2 999] (
			source = @world
			destination = {
				$p to @b
				remaining to @c
			}
		)`,
		`{
			"p": "3/2"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{
			Error: "portion must be",
		},
	)
}

func TestSource(t *testing.T) {
	testJSON(t,
		`vars {
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
)`,
		`{
			"balance": "users:001",
			"payment": "payments:001",
			"seller": "users:002"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"GEM": 3,
			},
			"payments:001": {
				"GEM": 12,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      3,
					Source:      "users:001",
					Destination: "users:002",
				},
				{
					Asset:       "GEM",
					Amount:      12,
					Source:      "payments:001",
					Destination: "users:002",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestAllocation(t *testing.T) {
	testJSON(t,
		`vars {
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
)`,
		`{
			"rider": "users:001",
			"driver": "users:002"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"GEM": 15,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      13,
					Source:      "users:001",
					Destination: "users:002",
				},
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "users:001",
					Destination: "a",
				},
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "users:001",
					Destination: "b",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestDynamicAllocation(t *testing.T) {
	testJSON(t,
		`vars {
	portion $p
}
send [GEM 15] (
	source = @a
	destination = {
		80% to @b
		$p to @c
		remaining to @d
	}
)`,
		`{
			"p": "15%"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"GEM": 15,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      13,
					Source:      "a",
					Destination: "b",
				},
				{
					Asset:       "GEM",
					Amount:      2,
					Source:      "a",
					Destination: "c",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSendAll(t *testing.T) {
	testJSON(t,
		`send [USD/2 *] (
  source = @users:001
  destination = @platform
)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"USD/2": 17,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "USD/2",
					Amount:      17,
					Source:      "users:001",
					Destination: "platform",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSendAllMulti(t *testing.T) {
	testJSON(t,
		`send [USD/2 *] (
			source = {
			  @users:001:wallet
			  @users:001:credit
			}
			destination = @platform
		  )
		  `,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001:wallet": {
				"USD/2": 19,
			},
			"users:001:credit": {
				"USD/2": 22,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "USD/2",
					Amount:      19,
					Source:      "users:001:wallet",
					Destination: "platform",
				},
				{
					Asset:       "USD/2",
					Amount:      22,
					Source:      "users:001:credit",
					Destination: "platform",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestInsufficientFunds(t *testing.T) {
	testJSON(t,
		`vars {
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
)`,
		`{
			"balance": "users:001",
			"payment": "payments:001",
			"seller": "users:002"
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"GEM": 3,
			},
			"payments:001": {
				"GEM": 12,
			},
		},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_FAIL_INSUFFICIENT_FUNDS,
		},
	)
}

func TestWorldSource(t *testing.T) {
	testJSON(t,
		`send [GEM 15] (
			source = {
				@a
				@world
			}
			destination = @b
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"GEM": 1,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "a",
					Destination: "b",
				},
				{
					Asset:       "GEM",
					Amount:      14,
					Source:      "world",
					Destination: "b",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestNoEmptyPostings(t *testing.T) {
	testJSON(t,
		`send [GEM 2] (
			source = @world
			destination = {
				90% to @a
				10% to @b
			}
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      2,
					Source:      "world",
					Destination: "a",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestNoEmptyPostings2(t *testing.T) {
	testJSON(t,
		`send [GEM *] (
			source = @foo
			destination = @bar
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"foo": {
				"GEM": 0,
			},
		},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_OK,
		},
	)
}

func TestAllocateDontTakeTooMuch(t *testing.T) {
	testJSON(t,
		`send [CREDIT 200] (
			source = {
				@users:001
				@users:002
			}
			destination = {
				1/2 to @foo
				1/2 to @bar
			}
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"users:001": {
				"CREDIT": 100,
			},
			"users:002": {
				"CREDIT": 110,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "CREDIT",
					Amount:      100,
					Source:      "users:001",
					Destination: "foo",
				},
				{
					Asset:       "CREDIT",
					Amount:      100,
					Source:      "users:002",
					Destination: "bar",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestMetadata(t *testing.T) {
	commission, _ := core.NewPortionSpecific(*big.NewRat(125, 1000))
	testJSON(t,
		`vars {
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
		)`,
		`{
			"sale": "sales:042"
		}`,
		map[string]map[string]core.Value{
			"sales:042": {
				"seller": core.Account("users:053"),
			},
			"users:053": {
				"commission": *commission,
			},
		},
		map[string]map[string]uint64{
			"sales:042": {
				"EUR/2": 2500,
			},
			"users:053": {
				"EUR/2": 500,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      88,
					Source:      "sales:042",
					Destination: "users:053",
				},
				{
					Asset:       "EUR/2",
					Amount:      12,
					Source:      "sales:042",
					Destination: "platform",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestTrackBalances(t *testing.T) {
	testJSON(t,
		`
		send [COIN 50] (
			source = @world
			destination = @a
		)
		send [COIN 100] (
			source = @a
			destination = @b
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"COIN": 50,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      50,
					Source:      "world",
					Destination: "a",
				},
				{
					Asset:       "COIN",
					Amount:      100,
					Source:      "a",
					Destination: "b",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestTrackBalances2(t *testing.T) {
	testJSON(t,
		`
		send [COIN 50] (
			source = @a
			destination = @z
		)
		send [COIN 50] (
			source = @a
			destination = @z
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"COIN": 60,
			},
		},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_FAIL_INSUFFICIENT_FUNDS,
		},
	)
}

func TestTrackBalances3(t *testing.T) {
	testJSON(t,
		`send [COIN *] (
			source = @foo
			destination = {
				max [COIN 1000] to @bar
				remaining kept
			}
		)
		send [COIN *] (
			source = @foo
			destination = @bar
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"foo": {
				"COIN": 2000,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      1000,
					Source:      "foo",
					Destination: "bar",
				},

				{
					Asset:       "COIN",
					Amount:      1000,
					Source:      "foo",
					Destination: "bar",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSourceAllotment(t *testing.T) {
	testJSON(t,
		`
		send [COIN 100] (
			source = {
				60% from @a
				35.5% from @b
				4.5% from @c
			}
			destination = @d
		)
		`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"COIN": 100,
			},
			"b": {
				"COIN": 100,
			},
			"c": {
				"COIN": 100,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      61,
					Source:      "a",
					Destination: "d",
				},
				{
					Asset:       "COIN",
					Amount:      35,
					Source:      "b",
					Destination: "d",
				},
				{
					Asset:       "COIN",
					Amount:      4,
					Source:      "c",
					Destination: "d",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSourceOverlapping(t *testing.T) {
	testJSON(t,
		`
		send [COIN 99] (
			source = {
				15% from {
					@b
					@a
				}
				30% from @a
				remaining from @a
			}
			destination = @world
		)
		`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"COIN": 99,
			},
			"b": {
				"COIN": 3,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      3,
					Source:      "b",
					Destination: "world",
				},
				{
					Asset:       "COIN",
					Amount:      12,
					Source:      "a",
					Destination: "world",
				},
				{
					Asset:       "COIN",
					Amount:      30,
					Source:      "a",
					Destination: "world",
				},
				{
					Asset:       "COIN",
					Amount:      54,
					Source:      "a",
					Destination: "world",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestSourceComplex(t *testing.T) {
	testJSON(t,
		`
		vars {
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
		)
		`,
		`{
			"max": {
				"asset": "COIN",
				"amount": 120
			}
		}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"COIN": 1000,
			},
			"b": {
				"COIN": 40,
			},
			"c": {
				"COIN": 1000,
			},
			"d": {
				"COIN": 1000,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      4,
					Source:      "a",
					Destination: "platform",
				},
				{
					Asset:       "COIN",
					Amount:      40,
					Source:      "b",
					Destination: "platform",
				},
				{
					Asset:       "COIN",
					Amount:      56,
					Source:      "c",
					Destination: "platform",
				},
				{
					Asset:       "COIN",
					Amount:      100,
					Source:      "d",
					Destination: "platform",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestDestinationComplex(t *testing.T) {
	testJSON(t,
		`
		send [COIN 100] (
			source = @world
			destination = {
				20% to @a
				20% kept
				60% to {
					max [COIN 10] to @b
					remaining to @c
				}
			}
		)
		`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "COIN",
					Amount:      20,
					Source:      "world",
					Destination: "a",
				},
				{
					Asset:       "COIN",
					Amount:      10,
					Source:      "world",
					Destination: "b",
				},
				{
					Asset:       "COIN",
					Amount:      50,
					Source:      "world",
					Destination: "c",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
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

	m := NewMachine(p)

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

	m := NewMachine(p)

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
			if string(v) != string(expected) {
				fmt.Printf("%v\n", string(v))
				fmt.Printf("%v\n", string(expected))
				t.Fatalf("unexpected transaction metadata")
			}
		}
	}
}
