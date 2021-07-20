package program

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/numary/machine/core"
)

type Program struct {
	Constants      []core.Value
	Instructions   []byte
	Variables      map[string]VarInfo
	NeededBalances map[core.Address]map[core.Address]struct{}
}

type VarInfo struct {
	Ty   core.Type
	Addr core.Address
}

func (p Program) Print() {
	fmt.Println("Program:\nINSTRUCTIONS")
	for i := 0; i < len(p.Instructions); i++ {
		fmt.Printf("%02d----- ", i)
		switch p.Instructions[i] {
		case OP_APUSH:
			fmt.Print("OP_APUSH\n")
			address := binary.LittleEndian.Uint16(p.Instructions[i+1 : i+3])
			if address >= 32768 {
				fmt.Printf("%02d-%02d   #VAR(%d)\n", i+1, i+3, address-32768)
			} else {
				fmt.Printf("%02d-%02d   #CONST(%d)\n", i+1, i+3, address)
			}
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
		case OP_SOURCE:
			fmt.Print("OP_SOURCE\n")
		case OP_ALLOC:
			fmt.Print("OP_ALLOC\n")
		default:
			fmt.Print("Unknown opcode")
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

func (p *Program) ParseVariables(vars map[string]core.Value) ([]core.Value, error) {
	variables := make([]core.Value, len(p.Variables))
	if len(vars) != len(p.Variables) {
		return nil, fmt.Errorf(
			"mismatching number of variables: %v != %v",
			len(vars),
			len(p.Variables))
	}
	for name, info := range p.Variables {
		if val, ok := vars[name]; ok && val.GetType() == info.Ty {
			variables[info.Addr.ToIdx()] = val
		} else {
			return nil, fmt.Errorf("missing variables: %v", name)
		}
	}
	return variables, nil
}

func (p *Program) ParseVariablesJSON(vars map[string]json.RawMessage) ([]core.Value, error) {
	variables := make([]core.Value, len(p.Variables))
	for name, info := range p.Variables {
		if val, ok := vars[name]; ok {
			var value core.Value
			switch info.Ty {
			case core.TYPE_ACCOUNT:
				var account core.Account
				err := json.Unmarshal(val, &account)
				if err != nil {
					return nil, err
				}
				value = account
			case core.TYPE_ASSET:
				var asset core.Asset
				err := json.Unmarshal(val, &asset)
				if err != nil {
					return nil, err
				}
				value = asset
			case core.TYPE_NUMBER:
				var number core.Number
				err := json.Unmarshal(val, &number)
				if err != nil {
					return nil, err
				}
				value = number
			case core.TYPE_MONETARY:
				var mon struct {
					Asset  string `json:"asset"`
					Amount uint64 `json:"amount"`
				}
				err := json.Unmarshal(val, &mon)
				if err != nil {
					return nil, err
				}
				value = core.Monetary{
					Asset:  core.Asset(mon.Asset),
					Amount: core.NewAmountSpecific(mon.Amount),
				}
			}
			variables[info.Addr.ToIdx()] = value
		} else {
			return nil, fmt.Errorf("missing variable %v", name)
		}
	}
	return variables, nil
}
