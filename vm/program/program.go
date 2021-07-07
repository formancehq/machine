package program

import (
	"encoding/binary"
	"fmt"

	"github.com/numary/machine/core"
)

type Program struct {
	Constants    []core.Value
	Instructions []byte
	Variables    map[string]VarInfo
}

type VarInfo struct {
	Ty   core.Type
	Addr Address
}

func (p Program) Print() {
	fmt.Println("Program:\nINSTRUCTIONS")
	for i := 0; i < len(p.Instructions); i++ {
		fmt.Printf("%02d----- ", i)
		switch p.Instructions[i] {
		case OP_APUSH:
			fmt.Print("OP_APUSH\n")
			fmt.Printf("%02d-%02d   #%d\n", i+1, i+3, binary.LittleEndian.Uint16(p.Instructions[i+1:i+3]))
			i += 2
		case OP_IPUSH:
			fmt.Print("OP_IPUSH\n")
			fmt.Printf("%02d-%02d   %d\n", i+1, i+9, binary.LittleEndian.Uint64(p.Instructions[i+1:i+9]))
			i += 8
		case OP_IADD:
			fmt.Print("OP_IADD\n")
		case OP_ISUB:
			fmt.Print("OP_ISUB\n")
		case OP_PRINT:
			fmt.Print("OP_PRINT\n")
		case OP_FAIL:
			fmt.Print("OP_FAIL\n")
		case OP_SEND:
			fmt.Print("OP_SEND\n")
		}
	}

	fmt.Println("CONSTANTS")
	i := 0
	for i = 0; i < len(p.Constants); i++ {
		fmt.Printf("%02d ", i)
		fmt.Printf("%s\n", p.Constants[i])
	}

	fmt.Println("VARIABLES")
	for name, info := range p.Variables {
		fmt.Printf("%02d ", info.Addr.ToIdx())
		fmt.Printf("%-4s\n", name)
	}
}
