package main

import (
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {

	p := compiler.Compile(`29 + 15 - 2`)

	machine := vm.NewMachine(p)

	machine.Program.Print()
	machine.Execute()
}
