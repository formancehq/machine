package main

import (
	"encoding/json"
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	p, err := compiler.Compile(`vars {
			account $rider
			account $driver
			monetary $value
		}
		send(value=$value, source=$rider, destination=$driver)`)

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

	exit_code, err := machine.ExecuteFromJSON(vars)
	if err != nil {
		panic(err)
	}
	fmt.Println("EXIT_CODE:", exit_code)
	fmt.Println(machine.Postings)
}
