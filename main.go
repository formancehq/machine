package main

import (
	"encoding/json"
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	p, err := compiler.Compile(`vars {
		account $balance
		account $payment
		account $seller
	}
	send [GEM 15] (
		source = {
			$balance
			$payment
		}
		destination = $seller
	)`)

	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage

	json.Unmarshal([]byte(`{
		"balance": "users:001",
		"payment": "payments:001",
		"seller": "users:002"
	}`), &vars)

	err = machine.SetVarsFromJSON(vars)
	if err != nil {
		panic(err)
	}
	machine.SetBalances(map[string]map[string]uint64{
		"users:001": {
			"GEM": 15,
		},
		"payments:001": {
			"GEM": 0,
		},
	})
	exit_code, err := machine.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
