package main

import (
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {

	p := compiler.Compile(`calc 29 + 15 - 2
fail`)

	machine := vm.NewMachine(p)

	machine.Program.Print()
	exit_code := machine.Execute()
	fmt.Println(exit_code)
}
