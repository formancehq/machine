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
	program, err := compiler.Compile(`
	send [GEM 40] (
		source = {
			50% from @foo allowing overdraft up to [GEM 10]
			50% from @bar
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
		balances := map[string]map[string]int64{
			"foo": {
				"GEM": 10,
			},
			"bar": {
				"GEM": 20,
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
