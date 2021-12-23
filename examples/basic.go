package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	program, err := compiler.Compile(`
	vars {
		boolean $cond0
		boolean $cond1
	}
	if $cond0 {
		set_tx_meta("first", 1)
		if $cond1 {
			set_tx_meta("second", 1)
		}
		set_tx_meta("first-end", 1)
	}
	set_tx_meta("end", 1)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Print(program)

	m := vm.NewMachine(program)

	m.SetVars(map[string]core.Value{
		"cond0": core.Boolean(true),
		"cond1": core.Boolean(true),
	})

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
	fmt.Println(m.TxMeta)
}
