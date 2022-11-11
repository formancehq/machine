package main

import (
	"fmt"

	"github.com/formancehq/machine/core"
	"github.com/formancehq/machine/script/compiler"
	"github.com/formancehq/machine/vm"
)

func main() {
	program, err := compiler.Compile(`
	send [COIN 99] (
		source = {
			15% from {
				@a
				@b
			}
			remaining from @a
		}
		destination = @world
	)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Print(program)

	m := vm.NewMachine(*program)

	_ = m.SetVars(map[string]core.Value{})

	{
		ch, err := m.ResolveResources()
		if err != nil {
			panic(err)
		}
		for range ch {

		}
	}

	{
		balances := map[string]map[string]*core.MonetaryInt{
			"a": {
				"COIN": core.NewMonetaryInt(500000),
			},
			"b": {
				"COIN": core.NewMonetaryInt(3500000),
			},
		}

		ch, err := m.ResolveBalances()
		if err != nil {
			panic(err)
		}
		for req := range ch {
			val := balances[req.Account][req.Asset]
			if req.Error != nil {
				panic(req.Error)
			}
			req.Response <- val
		}
	}

	m.Debug = true
	exitCode, err := m.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("Exit code:", exitCode)
	fmt.Println(m.Postings)
	fmt.Println(m.TxMeta)
}
