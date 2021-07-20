package vm

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	ledger "github.com/numary/ledger/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
)

type CaseResult struct {
	Printed  []core.Value
	Postings []ledger.Posting
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

func test(t *testing.T, code string, variables map[string]core.Value, balances map[string]map[string]uint64, expected CaseResult) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		err := m.SetVars(variables)
		if err != nil {
			return 0, err
		}
		err = m.SetBalances(balances)
		if err != nil {
			return 0, err
		}
		return m.Execute()
	})
}

func testJSON(t *testing.T, code string, variables string, balances map[string]map[string]uint64, expected CaseResult) {
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
		err = m.SetBalances(balances)
		if err != nil {
			return 0, err
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
		}
		send [EUR/2 999] (
			source=$rider
			destination=$driver
		)`,
		`{
			"rider": "users:001",
			"driver": "users:002"
		}`,
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

func TestSendAll(t *testing.T) {
	testJSON(t,
		`send [USD/2 *] (
  source = @users:001
  destination = @platform
)`,
		`{}`,
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
			ExitCode: EXIT_FAIL,
		},
	)
}

func TestMissingBalance(t *testing.T) {
	testJSON(t,
		`send [GEM 15] (
			source = @a
			destination = @a
		)`,
		`{}`,
		map[string]map[string]uint64{
			"users:001": {
				"GEM": 3,
			},
			"payments:001": {
				"USD/2": 564,
			},
		},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: 0,
			Error:    "missing balance",
		},
	)
}

func TestMissingWorldBalance(t *testing.T) {
	testJSON(t,
		`send [GEM 15] (
			source = @world
			destination = @a
		)`,
		`{}`,
		map[string]map[string]uint64{},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      15,
					Source:      "world",
					Destination: "a",
				},
			},
			ExitCode: 1,
			Error:    "",
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
			ExitCode: 1,
			Error:    "",
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
			ExitCode: 1,
			Error:    "",
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
		map[string]map[string]uint64{
			"foo": {
				"GEM": 0,
			},
		},
		CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: 1,
			Error:    "",
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
		map[string]map[string]uint64{
			"users:001": {
				"CREDIT": 100,
			},
			"users:002": {
				"CREDIT": 100,
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
			ExitCode: 1,
			Error:    "",
		},
	)
}

func TestGetNeededBalances(t *testing.T) {
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
		t.Errorf("did not expect error on Compile, got: %v", err)
		return
	}

	m := NewMachine(p)

	err = m.SetVars(map[string]core.Value{
		"a": core.Account("a"),
	})
	if err != nil {
		t.Errorf("did not expect error on SetVars, got: %v", err)
		return
	}

	expected := map[string]map[string]struct{}{
		"a": {
			"GEM": {},
		},
		"b": {
			"GEM": {},
		},
	}
	actual, err := m.GetNeededBalances()
	if err != nil {
		t.Errorf("did not expect error on GetNeededBalances, got: %v", err)
		return
	}
	if !cmp.Equal(actual, expected) {
		t.Errorf("unexpected needed balances: %v", actual)
	}
}
