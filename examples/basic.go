package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"

	ledger "github.com/numary/ledger/pkg/core"
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
		balances := map[string]map[string]ledger.MonetaryInt{
			"a": {
				"COIN": *ledger.NewMonetaryInt(500000),
			},
			"b": {
				"COIN": *ledger.NewMonetaryInt(3500000),
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
	exit_code, err := m.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("Exit code:", exit_code)
	fmt.Println(m.Postings)
	fmt.Println(m.TxMeta)
}
