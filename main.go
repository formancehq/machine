package main

import (
	"encoding/json"
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {

	// 	p, err := compiler.Compile(`vars {
	// 		account $rider
	// }
	// print 29 + 15 - 2
	// send(sum = [EUR 100], source = $rider, destination = @bank)
	// fail`)

	p, err := compiler.Compile(`vars {
			account $rider
			account $driver
			monetary $value
		}
		send(value=$value, source=$rider, destination=$driver)`)

	// fmt.Println(p)

	if err != nil {
		panic(err)
	}

	p.Print()

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage

	json.Unmarshal([]byte(`{
		"rider":  "user:001",
		"driver": "user:002",
		"value": {
			"asset":  "GEM",
			"amount": 32
		}
	}`), &vars)

	fmt.Println("HELLO", vars)

	exit_code, err := machine.ExecuteFromJSON(vars)
	if err != nil {
		panic(err)
	}
	fmt.Println(exit_code)
	fmt.Println(machine.Postings)

	// vars {
	// 	account $rider
	// 	account $driver
	// }
	// send(sum=[EUR/2 999], source=$rider, destination=$driver)
}
