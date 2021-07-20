package main

import (
	"encoding/json"
	"fmt"

	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

func main() {
	p, err := compiler.Compile(`send [GEM 15] (
		source = {
			@a
			@b
		}
		destination = {
			80% to @c
			 8% to @d
			12% to @e
		}
	)`)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	machine := vm.NewMachine(p)

	var vars map[string]json.RawMessage
	json.Unmarshal([]byte(`{}`), &vars)
	err = machine.SetVarsFromJSON(vars)

	if err != nil {
		panic(err)
	}
	err = machine.SetBalances(map[string]map[string]uint64{
		"a": {
			"GEM": 3,
		},
		"b": {
			"GEM": 25,
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
