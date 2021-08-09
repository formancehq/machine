package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	p, err := compiler.Compile(`
	vars {
		account $sale
		account $seller = meta($sale, "seller")
		portion $commission = meta($seller, "commission")
	}
	send [EUR/2 53] (
		source = {
			@a
			@d
		}
		destination = {
			remaining to $seller
			$commission to @platform
		}
	)
	`)

	// p, err := compiler.Compile(`
	// vars {
	// 	account $sale
	// 	account $seller = meta($sale, "seller")
	// 	portion $commission = meta($seller, "commission")
	// }
	// send [EUR/2 53] (
	// 	source = {
	// 		50% from {
	// 			max [COIN 4] from @a
	// 			@b
	// 			@c
	// 		}
	// 		remaining from max [COIN 120] from @d
	// 	}
	// 	destination = {
	// 		remaining to $seller
	// 		$commission to @platform
	// 	}
	// )
	// `)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"sale": "sales:042"
	}`), &vars)
	err = machine.SetVarsFromJSON(vars)
	if err != nil {
		panic(err)
	}

	commission, err := core.NewPortionSpecific(*big.NewRat(125, 1000))
	if err != nil {
		panic(err)
	}
	meta := map[string]map[string]core.Value{
		"sales:042": {
			"seller": core.Account("users:053"),
		},
		"users:053": {
			"commission": *commission,
		},
	}
	balances := map[string]map[string]uint64{
		"sales:042": {
			"EUR/2": 2500,
		},
		"users:053": {
			"EUR/2": 500,
		},
	}

	{
		ch, err := machine.ResolveResources()
		if err != nil {
			panic(err)
		}
		for req := range ch {
			fmt.Println(req)
			val := meta[req.Account][req.Key]
			if req.Error != nil {
				panic(req.Error)
			}
			req.Response <- val
		}
	}

	{
		ch, err := machine.ResolveBalances()
		if err != nil {
			panic(err)
		}
		for req := range ch {
			fmt.Println(req)
			val := balances[req.Account][req.Asset]
			if req.Error != nil {
				panic(req.Error)
			}
			req.Response <- val
		}
	}
	exit_code, err := machine.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
