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

	exit_code, err := machine.ExecuteFromJSON(vars, map[string]map[string]uint64{
		"users:001": {
			"GEM": 15,
		},
		"payments:001": {
			"GEM": 0,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}

// func main() {
// 	p, err := compiler.Compile(`vars {
// 	account $wallet
// 	account $payment
// 	account $driver
// 	monetary $value
// }

// send $value (
// 	source = {
// 		$wallet
// 		$payment
// 	}
// 	destination = {
// 		80% to $driver
// 		2/25 to @bank
// 		12% to @bank2
// 	}
// )`)

// 	if err != nil {
// 		panic(err)
// 	}

// 	p.Print()

// 	machine := vm.NewMachine(p)

// 	var vars map[string]json.RawMessage

// 	json.Unmarshal([]byte(`{
// 		"wallet":  "users:001",
// 		"payment":  "payments:001",
// 		"driver": "users:002",
// 		"value": {
// 			"asset":  "GEM",
// 			"amount": 15
// 		}
// 	}`), &vars)

// 	exit_code, err := machine.ExecuteFromJSON(vars, map[string]map[string]uint64{
// 		"users:001": {
// 			"GEM": 3,
// 		},
// 		"payments:001": {
// 			"GEM": 12,
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("EXIT_CODE:", exit_code)
// 	fmt.Println(machine.Postings)
// }
