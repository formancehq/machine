package vm

import (
	"testing"

	"numary.com/machine/script/compiler"
)

func TestMachine(t *testing.T) {
	cases := []struct {
		Case     string
		Expected []byte
	}{
		{
			Case:     "29 + 15 - 2",
			Expected: []byte{42},
		},
		{
			Case:     "1 + 1",
			Expected: []byte{2},
		},
		{
			Case:     "1",
			Expected: []byte{1},
		},
	}

	for _, c := range cases {
		p := compiler.Compile(c.Case)

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
				t.Fail()
			}
		}
	}
}
