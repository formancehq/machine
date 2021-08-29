package main

import (
	"fmt"

	"github.com/numary/machine/core"
	"github.com/numary/machine/script/compiler"
	"github.com/numary/machine/vm"
)

/*

push @world             @world
push [COIN 100]         @world [COIN 100]
asset                   @world COIN
take_all                [COIN * @world]
push [COIN 100]         [COIN * @world] [COIN 100]
take                    [COIN * @world] [COIN 100 @world]
swap                    [COIN 100 @world] [COIN * @world]
repay                   [COIN 100 @world]
// dest allotment
push 50%                [COIN 100 @world] 50%
push 50%                [COIN 100 @world] 50% 50%
push 2                [COIN 100 @world] 50% 50% 2
make_allotment          [COIN 100 @world] { 50% : 50% }
allocate                [COIN 50 @world] [COIN 50 @world]
// to @a
push @a                 [COIN 50 @world] [COIN 50 @world] @a
send                    [COIN 50 @world]
// to {...}
// max
push [COIN 10]          [COIN 50 @world] [COIN 10]
take_max                [COIN 40 @world] [COIN 10 @world]
// to @b
push @b                 [COIN 40 @world] [COIN 10 @world] @b
send                    [COIN 40 @world]
// to @c
push @c
send
*/

func main() {
	program, err := compiler.Compile(`
	send [COIN 100] (
		source = @world
		destination = {
			50% to @a
			50% to {
				max [COIN 10] to @b
				@c
			}
		}
	)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Print(program)

	m := vm.NewMachine(program)

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
		ch, err := m.ResolveBalances()
		if err != nil {
			panic(err)
		}
		for range ch {

		}
	}

	m.Debug = true
	exit_code, err := m.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println("Exit code:", exit_code)
	fmt.Println(m.Postings)
}
