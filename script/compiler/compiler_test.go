package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

type CaseResult struct {
	Instructions []byte
	Constants    []core.Value
	Variables    []string
	Error        string
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := Compile(c.Case)

	if c.Expected.Error != "" {
		if err == nil {
			t.Error(errors.New("expected error and got none"))
			return
		} else if !strings.Contains(err.Error(), c.Expected.Error) {
			t.Error(fmt.Errorf("error is not the one expected: %v", err))
			return
		}
	} else {
		if err != nil {
			t.Error(fmt.Errorf("did not expect error: %v", err))
			return
		} else if p == nil {
			t.Error(errors.New("did not receive any output"))
			return
		} else if !bytes.Equal(p.Instructions, c.Expected.Instructions) {
			t.Error(fmt.Errorf("generated program is incorrect: %v", *p))
			fmt.Println(p.Instructions, "vs", c.Expected.Instructions)
			return
		} else if len(p.Constants) != len(c.Expected.Constants) {
			t.Error(fmt.Errorf("unexpected program constants: %v", *p))
			return
		} else {
			for i := range c.Expected.Constants {
				if !core.ValueEquals(p.Constants[i], c.Expected.Constants[i]) {
					t.Error(fmt.Errorf("unexpected program constants: %v", *p))
					return
				}
			}
		}
	}
}

func TestSimplePrint(t *testing.T) {
	test(t, TestCase{
		Case: "print 1",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PRINT,
			},
			Constants: []core.Value{},
			Error:     "",
		},
	})
}

func TestCompositeExpr(t *testing.T) {
	test(t, TestCase{
		Case: "print 29 + 15 - 2",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 29, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IPUSH, 15, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IADD,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_ISUB,
				program.OP_PRINT,
			},
			Constants: []core.Value{},
			Error:     "",
		},
	})
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Instructions: []byte{program.OP_FAIL},
			Constants:    []core.Value{},
			Error:        "",
		},
	})
}

func TestConstant(t *testing.T) {
	user := core.Account("user:001")
	test(t, TestCase{
		Case: "print @user:001",
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Constants:    []core.Value{user},
			Error:        "",
		},
	})
}

func TestAllocationFractions(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 43] (
			source = @foo
			destination = {
				1/8 to @bar
				7/8 to @baz
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{
					Asset:  "EUR/2",
					Amount: core.NewAmountSpecific(43),
				},
				core.Account("foo"),
				core.Allotment{*big.NewRat(1, 8), *big.NewRat(7, 8)},
				core.Account("bar"),
				core.Account("baz"),
			},
			Error: "",
		},
	})
}

func TestAllocationPercentages(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 43] (
			source = @foo
			destination = {
				12.5% to @bar
				37.5% to @baz
				50% to @qux
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Constants: []core.Value{
				core.Monetary{
					Asset:  "EUR/2",
					Amount: core.NewAmountSpecific(43),
				},
				core.Account("foo"),
				core.Allotment{*big.NewRat(125, 1000), *big.NewRat(375, 1000), *big.NewRat(1, 2)},
				core.Account("bar"),
				core.Account("baz"),
				core.Account("qux"),
			},
			Error: "",
		},
	})
}

func TestSend(t *testing.T) {
	alice := core.Account("alice")
	bob := core.Account("bob")
	test(t, TestCase{
		Case: `send [EUR/2 99] (
	source = @alice
	destination = @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Constants: []core.Value{core.Monetary{Asset: "EUR/2", Amount: core.NewAmountSpecific(99)}, alice, bob},
			Error: "",
		},
	})
}

func TestSendAll(t *testing.T) {
	alice := core.Account("alice")
	bob := core.Account("bob")
	test(t, TestCase{
		Case: `send [EUR/2 *] (
	source = @alice
	destination = @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Constants: []core.Value{core.Monetary{Asset: "EUR/2", Amount: core.NewAmountAll()}, alice, bob},
			Error: "",
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case: "print fail",
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "mismatched input",
		},
	})
}

func TestLogicError(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 200] (
			source = 200
			destination = @bob
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "expected",
		},
	})
}

func TestPreventTakeAllFromWorld(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM *] (
			source = @world
			destination = @foo
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Constants:    nil,
			Error:        "cannot",
		},
	})
}

// func TestTooManyConstants(t *testing.T) {
// 	script := "print 1"
// 	for i := 0; i < 20000; i++ {
// 		script += fmt.Sprintf("\nsend(monetary=[A%d 0], source=a%d, destination=b%d)", i, i, i)
// 	}
// 	test(t, TestCase{
// 		Case: script,
// 		Expected: CaseResult{
// 			Instructions: nil,
// 			Constants:    nil,
// 			Error:        "exceeded",
// 		},
// 	})
// }
