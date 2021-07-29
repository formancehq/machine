package core

import (
	"encoding/json"
	"errors"
)

func TypenameToType(name string) Type {
	switch name {
	case "account":
		return TYPE_ACCOUNT
	case "portion":
		return TYPE_PORTION
	case "monetary":
		return TYPE_MONETARY
	default:
		return 0
	}
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
			Asset  string `json:"asset"`
			Amount uint64 `json:"amount"`
		}
		err := json.Unmarshal(data, &mon)
		if err != nil {
			return nil, err
		}
		value = Monetary{
			Asset:  Asset(mon.Asset),
			Amount: NewAmountSpecific(mon.Amount),
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
	default:
		return nil, errors.New("invalid type")
	}
	return &value, nil
}
