package program

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/numary/machine/core"
)

type Program struct {
	Instructions   []byte
	Resources      []Resource
	Parameters     map[string]core.Address
	NeededBalances map[core.Address]map[core.Address]struct{}
}

func (p Program) String() string {
	out := "Program:\nINSTRUCTIONS"
	for i := 0; i < len(p.Instructions); i++ {
		out += fmt.Sprintf("%02d----- ", i)
		switch p.Instructions[i] {
		case OP_APUSH:
			out += "OP_APUSH\n"
			address := binary.LittleEndian.Uint16(p.Instructions[i+1 : i+3])
			if address >= 32768 {
				out += fmt.Sprintf("%02d-%02d   #VAR(%d)\n", i+1, i+3, address-32768)
			} else {
				out += fmt.Sprintf("%02d-%02d   #CONST(%d)\n", i+1, i+3, address)
			}
			i += 2
		case OP_IPUSH:
			out += "OP_IPUSH\n"
			out += fmt.Sprintf("%02d-%02d   %d\n", i+1, i+9, binary.LittleEndian.Uint64(p.Instructions[i+1:i+9]))
			i += 8
		case OP_IADD:
			out += "OP_IADD\n"
		case OP_ISUB:
			out += "OP_ISUB\n"
		case OP_PRINT:
			out += "OP_PRINT\n"
		case OP_FAIL:
			out += "OP_FAIL\n"
		case OP_SEND:
			out += "OP_SEND\n"
		case OP_SOURCE:
			out += "OP_SOURCE\n"
		case OP_ALLOC:
			out += "OP_ALLOC\n"
		case OP_MAKE_ALLOTMENT:
			out += "OP_MAKE_ALLOTMENT\n"
		default:
			out += "Unknown opcode"
		}
	}

	out += fmt.Sprintln("CONSTANTS")
	i := 0
	for i = 0; i < len(p.Resources); i++ {
		out += fmt.Sprintf("%02d ", i)
		out += fmt.Sprintf("%v\n", p.Resources[i])
	}
	return out
}

// func (p *Program) ParseVariables(vars map[string]core.Value) ([]core.Value, error) {
// 	variables := make([]core.Value, len(p.Resources))
// 	if len(vars) != len(p.Variables) {
// 		return nil, fmt.Errorf(
// 			"mismatching number of variables: %v != %v",
// 			len(vars),
// 			len(p.Variables))
// 	}
// 	for name, info := range p.Variables {
// 		if val, ok := vars[name]; ok && val.GetType() == info.Ty {
// 			variables[info.Addr.ToIdx()] = val
// 		} else {
// 			return nil, fmt.Errorf("missing variables: %v", name)
// 		}
// 	}
// 	return variables, nil
// }

func (p *Program) ParseVariables(vars map[string]core.Value) (map[string]core.Value, error) {
	variables := make(map[string]core.Value)
	for _, res := range p.Resources {
		if param, ok := res.(Parameter); ok {
			if val, ok := vars[param.Name]; ok && val.GetType() == param.Typ {
				variables[param.Name] = val
				delete(vars, param.Name)
			} else {
				return nil, fmt.Errorf("missing variables: %q", param.Name)
			}
		}
	}
	for name := range vars {
		return nil, fmt.Errorf("extraneous variable: %q", name)
	}
	return variables, nil
}

func (p *Program) ParseVariablesJSON(vars map[string]json.RawMessage) (map[string]core.Value, error) {
	variables := make(map[string]core.Value)
	for _, res := range p.Resources {
		if param, ok := res.(Parameter); ok {
			data, ok := vars[param.Name]
			if !ok {
				return nil, fmt.Errorf("missing variable: %q", param.Name)
			}
			var value core.Value
			switch param.Typ {
			case core.TYPE_ACCOUNT:
				var account core.Account
				err := json.Unmarshal(data, &account)
				if err != nil {
					return nil, err
				}
				value = account
			case core.TYPE_ASSET:
				var asset core.Asset
				err := json.Unmarshal(data, &asset)
				if err != nil {
					return nil, err
				}
				value = asset
			case core.TYPE_NUMBER:
				var number core.Number
				err := json.Unmarshal(data, &number)
				if err != nil {
					return nil, err
				}
				value = number
			case core.TYPE_MONETARY:
				var mon struct {
					Asset  string `json:"asset"`
					Amount uint64 `json:"amount"`
				}
				err := json.Unmarshal(data, &mon)
				if err != nil {
					return nil, err
				}
				value = core.Monetary{
					Asset:  core.Asset(mon.Asset),
					Amount: core.NewAmountSpecific(mon.Amount),
				}
			case core.TYPE_PORTION:
				var s string
				err := json.Unmarshal(data, &s)
				if err != nil {
					return nil, err
				}
				res, err := core.ParsePortionSpecific(s)
				if err != nil {
					return nil, err
				}
				value = *res
			default:
				return nil, errors.New("unexpected variable type in program")
			}
			variables[param.Name] = value
			delete(vars, param.Name)
		}
	}
	for name := range vars {
		return nil, fmt.Errorf("extraneous variable: %q", name)
	}
	return variables, nil
}
