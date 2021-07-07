package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

// print 29 + 15 - 2
// send(sum=[EUR 100], source = rider, destination = driver)
// fail

func main() {

	p, err := compiler.Compile(`vars {
		account $rider
}
print 29 + 15 - 2
send(sum = [EUR 100], source = $rider, destination = @bank)
fail`)

	// fmt.Println(p)

	if err != nil {
		panic(err)
	}

	machine := vm.NewMachine(p)

	acc := core.Account("user:001")

	machine.Program.Print()
	exit_code := machine.Execute(map[string]core.Value{
		"rider": &acc,
	})
	fmt.Println(exit_code)
	fmt.Println(machine.Postings)
}
