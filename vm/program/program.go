package program

import (
	"encoding/binary"
	"encoding/json"
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
	out := "Program:\nINSTRUCTIONS\n"
	for i := 0; i < len(p.Instructions); i++ {
		out += fmt.Sprintf("%02d----- ", i)
		switch p.Instructions[i] {
		case OP_APUSH:
			out += "OP_APUSH "
			address := binary.LittleEndian.Uint16(p.Instructions[i+1 : i+3])
			out += fmt.Sprintf("#%d\n", address)
			i += 2
		case OP_IPUSH:
			out += "OP_IPUSH "
			out += fmt.Sprintf("%d\n", binary.LittleEndian.Uint64(p.Instructions[i+1:i+9]))
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
		case OP_ALLOC:
			out += "OP_ALLOC\n"
		case OP_MAKE_ALLOTMENT:
			out += "OP_MAKE_ALLOTMENT\n"
		case OP_TAKE_ALL:
			out += "OP_TAKE_ALL\n"
		case OP_TAKE:
			out += "OP_TAKE\n"
		case OP_TAKE_SPLIT:
			out += "OP_TAKE_SPLIT\n"
		case OP_ASSEMBLE:
			out += "OP_ASSEMBLE\n"
		case OP_ASSET:
			out += "OP_ASSET\n"
		default:
			out += "Unknown opcode"
		}
	}

	out += fmt.Sprintln("RESOURCES")
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
			value, err := core.NewValueFromJSON(param.Typ, data)
			if err != nil {
				return nil, fmt.Errorf("invalid json for variable of %v of type %v: %v", param.Name, param.Typ, err)
			}
			variables[param.Name] = *value
			delete(vars, param.Name)
		}
	}
	for name := range vars {
		return nil, fmt.Errorf("extraneous variable: %q", name)
	}
	return variables, nil
}
