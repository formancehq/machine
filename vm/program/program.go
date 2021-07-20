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

func (p Program) String() string {
	out := "Program:\nINSTRUCTIONS"
	for i := 0; i < len(p.Instructions); i++ {
		out += fmt.Sprintf("%02d----- ", i)
		switch p.Instructions[i] {
		case OP_APUSH:
			out += fmt.Sprint("OP_APUSH\n")
			address := binary.LittleEndian.Uint16(p.Instructions[i+1 : i+3])
			if address >= 32768 {
				out += fmt.Sprintf("%02d-%02d   #VAR(%d)\n", i+1, i+3, address-32768)
			} else {
				out += fmt.Sprintf("%02d-%02d   #CONST(%d)\n", i+1, i+3, address)
			}
			i += 2
		case OP_IPUSH:
			out += fmt.Sprint("OP_IPUSH\n")
			out += fmt.Sprintf("%02d-%02d   %d\n", i+1, i+9, binary.LittleEndian.Uint64(p.Instructions[i+1:i+9]))
			i += 8
		case OP_IADD:
			out += fmt.Sprint("OP_IADD\n")
		case OP_ISUB:
			out += fmt.Sprint("OP_ISUB\n")
		case OP_PRINT:
			out += fmt.Sprint("OP_PRINT\n")
		case OP_FAIL:
			out += fmt.Sprint("OP_FAIL\n")
		case OP_SEND:
			out += fmt.Sprint("OP_SEND\n")
		case OP_SOURCE:
			out += fmt.Sprint("OP_SOURCE\n")
		case OP_ALLOC:
			out += fmt.Sprint("OP_ALLOC\n")
		case OP_MAKE_ALLOTMENT:
			out += fmt.Sprint("OP_MAKE_ALLOTMENT\n")
		default:
			out += fmt.Sprint("Unknown opcode")
		}
	}

	out += fmt.Sprintln("CONSTANTS")
	i := 0
	for i = 0; i < len(p.Constants); i++ {
		out += fmt.Sprintf("%02d ", i)
		out += fmt.Sprintf("%s\n", p.Constants[i])
	}

	out += fmt.Sprintln("VARIABLES")
	for name, info := range p.Variables {
		out += fmt.Sprintf("%02d ", info.Addr.ToIdx())
		out += fmt.Sprintf("%-4s\n", name)
	}
	return out
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
