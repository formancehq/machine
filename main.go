package main

import (
	"numary.com/machine/script/compiler"
	"numary.com/machine/vm"
)

func main() {

	p := compiler.Compile(`29 + 15 - 2`)

	machine := vm.NewMachine(p)

	machine.Program.Print()
	machine.Execute()
}
