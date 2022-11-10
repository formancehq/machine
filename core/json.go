package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ValueJSON struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func TypenameToType(name string) (Type, bool) {
	switch name {
	case "account":
		return TypeAccount, true
	case "asset":
		return TypeAsset, true
	case "number":
		return TypeNumber, true
	case "portion":
		return TypePortion, true
	case "monetary":
		return TypeMonetary, true
	default:
		return 0, false
	}
}

func NewValueFromTypedJSON(rawInput json.RawMessage) (*Value, error) {
	var input ValueJSON

	err := json.Unmarshal(rawInput, &input)
	if err != nil {
		return nil, err
	}

	typ, ok := TypenameToType(input.Type)
	if !ok {
		return nil, fmt.Errorf("unknown type: %v", input.Type)
	}

	return NewValueFromJSON(typ, input.Value)
}

func NewValueFromJSON(typ Type, data json.RawMessage) (*Value, error) {
	var value Value
	switch typ {
	case TypeAccount:
		var account Account
		err := json.Unmarshal(data, &account)
		if err != nil {
			return nil, err
		}
		value = account
	case TypeAsset:
		var asset Asset
		err := json.Unmarshal(data, &asset)
		if err != nil {
			return nil, err
		}
		value = asset
	case TypeNumber:
		var number Number
		err := json.Unmarshal(data, &number)
		if err != nil {
			return nil, err
		}
		value = number
	case TypeMonetary:
		var mon struct {
			Asset  string      `json:"asset"`
			Amount MonetaryInt `json:"amount"`
		}
		err := json.Unmarshal(data, &mon)
		if err != nil {
			return nil, err
		}
		value = Monetary{
			Asset:  Asset(mon.Asset),
			Amount: mon.Amount,
		}
	case TypePortion:
		var s string
		err := json.Unmarshal(data, &s)
		if err != nil {
			return nil, err
		}
		res, err := ParsePortionSpecific(s)
		if err != nil {
			return nil, err
		}
		value = *res
	case TypeString:
		var s String
		err := json.Unmarshal(data, &s)
		if err != nil {
			return nil, err
		}
		value = s
	default:
		return nil, errors.New("invalid type")
	}
	return &value, nil
}
