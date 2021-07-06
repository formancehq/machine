package program

import "fmt"

type Program struct {
	Data         []string
	Instructions []byte
	Vars         []string
}

func (p Program) Print() {
	fmt.Println("Program:\nINSTRUCTIONS")
	for i := 0; i < len(p.Instructions); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%08b\n", p.Instructions[i])
	}

	fmt.Println("RESOURCES")
	i := 0
	for i = 0; i < len(p.Data); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%s\n", p.Data[i])
	}

	for addr, name := range p.Vars {
		fmt.Printf("%-16d ", addr)
		fmt.Printf("%s\n", name)
	}
}
