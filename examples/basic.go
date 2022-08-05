package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

// send [GEM 50] (
// 	source = {
// 		@foo
// 		@bar
// 		@baz
// 	}
// 	destination = {
// 		50% to {
// 			max [GEM 4] to {
// 				50% kept
// 				25% to @arst
// 				25% kept
// 			}
// 			remaining to @thing
// 		}
// 		25% to @qux
// 		remaining to @quz
// 	}
// )

func main() {
	program, err := compiler.Compile(`send [COIN 100] (
		source = @world
		destination = {
			20% to @a
			20% kept
			60% to {
				max [COIN 10] to @b
				remaining to @c
			}
		}
	)`)
	if err != nil {
		panic(err)
	}
	fmt.Print(program)
	fmt.Print(program.NeededBalances)

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
		balances := map[string]map[string]int64{}

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
