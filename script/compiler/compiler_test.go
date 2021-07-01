package compiler

import (
	"fmt"
	"testing"

	"github.com/numary/machine/vm/program"
)

type CaseResult struct {
	Program []byte
	Error   error
}

func TestCompiler(t *testing.T) {
	cases := []struct {
		Case     string
		Expected CaseResult
	}{
		{
			Case: "calc 29 + 15 - 2",
			Expected: CaseResult{
				Program: []byte{program.OP_IPUSH, 29, program.OP_IPUSH, 15, program.OP_IADD, program.OP_IPUSH, 2, program.OP_ISUB, program.OP_PRINT},
				Error:   nil,
			},
		},
		{
			Case: "calc 1 + 1",
			Expected: CaseResult{
				Program: []byte{program.OP_IPUSH, 1, program.OP_IPUSH, 1, program.OP_IADD, program.OP_PRINT},
				Error:   nil,
			},
		},
		{
			Case: "calc 1",
			Expected: CaseResult{
				Program: []byte{program.OP_IPUSH, 1, program.OP_PRINT},
				Error:   nil,
			},
		},
		{
			Case: "fail",
			Expected: CaseResult{
				Program: []byte{program.OP_FAIL},
				Error:   nil,
			},
		},
		{
			Case: "calc fail",
			Expected: CaseResult{
				Program: nil,
				Error: &CompileError{
					SyntaxError{
						line:   1,
						column: 5,
						msg:    "no viable alternative at input 'calcfail'",
					},
				},
			},
		},
	}

	for _, c := range cases {
		p, err := Compile(c.Case)

		if err != nil {
			if c.Expected.Error == nil {
				t.Error(fmt.Errorf("did not expect error: %v", err))
			}
			if err.Error() != c.Expected.Error.Error() {
				t.Error(fmt.Errorf("error is not the one expected: %v", err))
			}
		}

		if p != nil {
			for i, op := range p {
				if op != c.Expected.Program[i] {
					t.Error(fmt.Errorf("generated program is incorrect: %v", p))
				}
			}
		}
	}
}
