package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	program, err := compiler.Compile(`
	send [COIN 100] (
		source = @world
		destination = {
			50% to @a
			50% to {
				max [COIN 10] to @b
				@c
			}
		}
	)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Print(program)

	m := vm.NewMachine(program)

	m.SetVars(map[string]core.Value{})

	{
		ch, err := m.ResolveResources()
		if err != nil {
			panic(err)
		}
		for range ch {

		}
	}

	{
		ch, err := m.ResolveBalances()
		if err != nil {
			panic(err)
		}
		for range ch {

		}
	}

	m.Debug = true
	exit_code, err := m.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("Exit code:", exit_code)
	fmt.Println(m.Postings)
}
