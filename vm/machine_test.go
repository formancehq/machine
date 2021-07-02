package vm

import (
	"fmt"
	"testing"

	"github.com/numary/machine/script/compiler"
)

func TestMachine(t *testing.T) {
	cases := []struct {
		Case     string
		Expected []byte
	}{
		{
			Case:     "calc 29 + 15 - 2",
			Expected: []byte{42},
		},
		{
			Case:     "calc 1 + 1",
			Expected: []byte{2},
		},
		{
			Case:     "calc 1",
			Expected: []byte{1},
		},
	}

	for _, c := range cases {
		p, err := compiler.Compile(c.Case)

		if err != nil {
			t.Error(fmt.Errorf("compile error: %v", err))
		}

		printed := []byte{}

		machine := NewMachine(p)
		machine.Printer = func(c chan byte) {
			for v := range c {
				printed = append(printed, v)
			}
		}
		machine.Execute()

		for i := range printed {
			if printed[i] != c.Expected[i] {
				t.Error(fmt.Errorf("unexpected output: %v", printed[i]))
			}
		}
	}
}
