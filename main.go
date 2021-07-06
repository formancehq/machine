package main

import (
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {

	p, err := compiler.Compile(`print 29 + 15 - 2
send(monetary=[EUR 100], source=alice, destination=bob)
fail`)

	// fmt.Println(p)

	if err != nil {
		panic(err)
	}

	machine := vm.NewMachine(p)

	machine.Program.Print()
	exit_code := machine.Execute(map[string]string{})
	fmt.Println(exit_code)
	fmt.Println(machine.Postings)
}
