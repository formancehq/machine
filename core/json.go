package core

import (
	"encoding/json"
	"errors"
	"fmt"

	ledger "github.com/numary/ledger/pkg/core"
)

type ValueJSON struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func TypenameToType(name string) (Type, bool) {
	switch name {
	case "account":
		return TYPE_ACCOUNT, true
	case "asset":
		return TYPE_ASSET, true
	case "number":
		return TYPE_NUMBER, true
	case "portion":
		return TYPE_PORTION, true
	case "monetary":
		return TYPE_MONETARY, true
	default:
		return 0, false
	}
}

func NewValueFromTypedJSON(raw_input json.RawMessage) (*Value, error) {
	var input ValueJSON

	err := json.Unmarshal(raw_input, &input)
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
	case TYPE_ACCOUNT:
		var account Account
		err := json.Unmarshal(data, &account)
		if err != nil {
			return nil, err
		}
		value = account
	case TYPE_ASSET:
		var asset Asset
		err := json.Unmarshal(data, &asset)
		if err != nil {
			return nil, err
		}
		value = asset
	case TYPE_NUMBER:
		var number Number
		err := json.Unmarshal(data, &number)
		if err != nil {
			return nil, err
		}
		value = number
	case TYPE_MONETARY:
		var mon struct {
			Asset  string             `json:"asset"`
			Amount ledger.MonetaryInt `json:"amount"`
		}
		err := json.Unmarshal(data, &mon)
		if err != nil {
			return nil, err
		}
		value = Monetary{
			Asset:  Asset(mon.Asset),
			Amount: mon.Amount,
		}
	case TYPE_PORTION:
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
	case TYPE_STRING:
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
