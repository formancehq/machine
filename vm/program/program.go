package program

import "fmt"

type Program struct {
	Constants    []string
	Instructions []byte
	Variables    []string
}

func (p Program) Print() {
	fmt.Println("Program:\nINSTRUCTIONS")
	for i := 0; i < len(p.Instructions); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%08b\n", p.Instructions[i])
	}

	fmt.Println("CONSTANTS")
	i := 0
	for i = 0; i < len(p.Constants); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%s\n", p.Constants[i])
	}

	fmt.Println("VARIABLES")
	for addr, name := range p.Variables {
		fmt.Printf("%-16d ", addr)
		fmt.Printf("%s\n", name)
	}
}
