package vm

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	ledger "github.com/numary/ledger/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
)

type CaseResult struct {
	Printed  []core.Value
	Postings []ledger.Posting
	ExitCode byte
	Error    error
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

func test(t *testing.T, code string, variables map[string]core.Value, expected CaseResult) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		return m.Execute(variables)
	})
}

func testJSON(t *testing.T, code string, variables string, expected CaseResult) {
	testimpl(t, code, expected, func(m *Machine) (byte, error) {
		var v map[string]json.RawMessage
		err := json.Unmarshal([]byte(variables), &v)
		if err != nil {
			panic("test case was invalid")
		}
		return m.ExecuteFromJSON(v)
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

	if err != expected.Error {
		t.Error(fmt.Errorf("unexpected execution error: %v", err))
	}

	wg.Wait()

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
			"rider":  core.Account("user:001"),
			"driver": core.Account("user:002"),
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      999,
					Source:      "user:001",
					Destination: "user:002",
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
			"rider": "user:001",
			"driver": "user:002"
		}`,
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "EUR/2",
					Amount:      999,
					Source:      "user:001",
					Destination: "user:002",
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
			"rider": "user:001",
			"driver": "user:002"
		}`,
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      13,
					Source:      "user:001",
					Destination: "user:002",
				},
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "user:001",
					Destination: "a",
				},
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "user:001",
					Destination: "b",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}
