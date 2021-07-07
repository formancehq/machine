package vm

import (
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
}

type TestCase struct {
	Code      string
	Variables map[string]core.Value
	Expected  CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := compiler.Compile(c.Code)

	if err != nil {
		t.Error(fmt.Errorf("compile error: %v", err))
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
	exit_code := machine.Execute(c.Variables)

	wg.Wait()

	if exit_code != c.Expected.ExitCode {
		t.Error(fmt.Errorf("unexpected exit code: %v", exit_code))
		return
	}

	if len(machine.Postings) != len(c.Expected.Postings) {
		t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings))
		return
	} else {
		for i := range machine.Postings {
			if machine.Postings[i] != c.Expected.Postings[i] {
				t.Error(fmt.Errorf("unexpected postings output: %v", machine.Postings[i]))
				return
			}
		}
	}

	if len(printed) != len(c.Expected.Printed) {
		t.Error(fmt.Errorf("unexpected print output: %v", printed))
		return
	} else {
		for i := range printed {
			if printed[i] != c.Expected.Printed[i] {
				t.Error(fmt.Errorf("unexpected print output: %v", printed[i]))
				return
			}
		}
	}
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Code:      "fail",
		Variables: map[string]core.Value{},
		Expected: CaseResult{
			Printed:  []core.Value{},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_FAIL,
		},
	})
}

func TestPrint(t *testing.T) {
	test(t, TestCase{
		Code:      "print 29 + 15 - 2",
		Variables: map[string]core.Value{},
		Expected: CaseResult{
			Printed:  []core.Value{core.Number(42)},
			Postings: []ledger.Posting{},
			ExitCode: EXIT_OK,
		},
	})
}

func TestSend(t *testing.T) {
	test(t, TestCase{
		Code:      "send(sum=[EUR/2 100], source=@alice, destination=@bob)",
		Variables: map[string]core.Value{},
		Expected: CaseResult{
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
	})
}

func TestVariables(t *testing.T) {
	test(t, TestCase{
		Code: `vars {
			account $rider
			account $driver
		}
		send(sum=[EUR/2 999], source=$rider, destination=$driver)`,
		Variables: map[string]core.Value{
			"rider":  core.Account("user:001"),
			"driver": core.Account("user:002"),
		},
		Expected: CaseResult{
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
	})
}
