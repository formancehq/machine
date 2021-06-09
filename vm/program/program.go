package program

import "fmt"

type Program []byte

func (p Program) Print() {
	for i := 0; i < len(p); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%08b\n", p[i])
	}
}
