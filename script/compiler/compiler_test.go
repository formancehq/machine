package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/numary/machine/vm/program"
)

type CaseResult struct {
	Instructions []byte
	Data         []string
	Error        error
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := Compile(c.Case)

	if c.Expected.Error != nil {
		if err == nil {
			t.Error(errors.New("expected error and got none"))
			return
		} else if err.Error() != c.Expected.Error.Error() {
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
			t.Error(fmt.Errorf("generated program is incorrect: %v", p.Instructions))
			return
		} else if len(p.Data) != len(c.Expected.Data) {
			t.Error(fmt.Errorf("unexpected program data: %v", p.Data))
			return
		} else {
			for i := range c.Expected.Data {
				if p.Data[i] != c.Expected.Data[i] {
					t.Error(fmt.Errorf("unexpected program data: %v", p.Data))
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
				program.OP_PUSH8, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PRINT,
			},
			Data:  []string{},
			Error: nil,
		},
	})
}

func TestCompositeExpr(t *testing.T) {
	test(t, TestCase{
		Case: "print 29 + 15 - 2",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_PUSH8, 29, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PUSH8, 15, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IADD,
				program.OP_PUSH8, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_ISUB,
				program.OP_PRINT,
			},
			Data:  []string{},
			Error: nil,
		},
	})
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Instructions: []byte{program.OP_FAIL},
			Data:         []string{},
			Error:        nil,
		},
	})
}

func TestSend(t *testing.T) {
	test(t, TestCase{
		Case: "send(monetary=[EUR/2 99], source=alice, destination=bob)",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_PUSH2, 00, 00,
				program.OP_PUSH8, 99, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PUSH2, 01, 00,
				program.OP_PUSH2, 02, 00,
				program.OP_SEND,
			}, Data: []string{"EUR/2", "alice", "bob"},
			Error: nil,
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case: "print fail",
		Expected: CaseResult{
			Instructions: nil,
			Data:         nil,
			Error: &CompileError{
				SyntaxError{
					line:   1,
					column: 6,
					msg:    "mismatched input 'fail' expecting {'[', IDENTIFIER, NUMBER, ASSET}",
				},
			},
		},
	})
}
