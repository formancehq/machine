package main

import (
	"encoding/json"
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	p, err := compiler.Compile(`vars {
		account $user_wallet
		account $user_credit
	  }
	  
	  send [CREDIT 100] (
		source = @world
		destination = $user_wallet
	  )
	  
	  send [CREDIT 100] (
		source = @world
		destination = $user_credit
	  )
	  
	  send [CREDIT 200] (
		source = {
		  $user_wallet
		  $user_credit
		}
		destination = {
		  50% to @feebles
		  50% to @boofaz
		}
	  )`)

	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{
		"user_wallet": "users:020:wallet",
		"user_credit": "users:020:credit"
	}`), &vars)
	err = machine.SetVarsFromJSON(vars)

	if err != nil {
		panic(err)
	}
	err = machine.SetBalances(map[string]map[string]uint64{
		"users:020:credit": {
			"CREDIT": 0,
		},
		"users:020:wallet": {
			"CREDIT": 0,
		},
	})
	if err != nil {
		panic(err)
	}
	exit_code, err := machine.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
