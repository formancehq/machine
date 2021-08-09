package core

import (
	"fmt"
	"reflect"
)

type Type byte

const (
	TYPE_ACCOUNT   = Type(iota + 1) // address of an account
	TYPE_ASSET                      // name of an asset
	TYPE_NUMBER                     // 64bit unsigned integer
	TYPE_MONETARY                   // [asset number]
	TYPE_PORTION                    // rational number between 0 and 1 both exclusive
	TYPE_ALLOTMENT                  // list of portions
	TYPE_AMOUNT                     // either ALL or a SPECIFIC number
	TYPE_FUNDING                    // (asset, []{amount, account})
)

func (t Type) String() string {
	switch t {
	case TYPE_ACCOUNT:
		return "account"
	case TYPE_ASSET:
		return "asset"
	case TYPE_NUMBER:
		return "number"
	case TYPE_MONETARY:
		return "monetary"
	case TYPE_PORTION:
		return "portion"
	case TYPE_ALLOTMENT:
		return "allotment"
	case TYPE_AMOUNT:
		return "amount"
	default:
		return "invalid type"
	}
}

type Value interface {
	isValue()
	GetType() Type
}

type Account string

func (Account) isValue()      {}
func (Account) GetType() Type { return TYPE_ACCOUNT }
func (a Account) String() string {
	return fmt.Sprintf("@%v", string(a))
}

type Asset string

func (Asset) isValue()      {}
func (Asset) GetType() Type { return TYPE_ASSET }
func (a Asset) String() string {
	return fmt.Sprintf("%v", string(a))
}

type Number uint64

func (Number) isValue()      {}
func (Number) GetType() Type { return TYPE_NUMBER }
func (n Number) String() string {
	return fmt.Sprintf("%v", uint64(n))
}

type Monetary struct {
	Asset  Asset  `json:"asset"`
	Amount uint64 `json:"amount"`
}

func (a Monetary) String() string {
	return fmt.Sprintf("[%v %v]", a.Asset, a.Amount)
}

func (Monetary) isValue()      {}
func (Monetary) GetType() Type { return TYPE_MONETARY }

func (Allotment) isValue()      {}
func (Allotment) GetType() Type { return TYPE_ALLOTMENT }

func (Portion) isValue()      {}
func (Portion) GetType() Type { return TYPE_PORTION }

func (Funding) isValue()      {}
func (Funding) GetType() Type { return TYPE_FUNDING }

type HasAsset interface {
	GetAsset() Asset
}

func (a Asset) GetAsset() Asset    { return a }
func (m Monetary) GetAsset() Asset { return m.Asset }
func (f Funding) GetAsset() Asset  { return f.Asset }

func ValueEquals(lhs, rhs Value) bool {
	if reflect.TypeOf(lhs) != reflect.TypeOf(rhs) {
		return false
	}
	if lhsa, ok := lhs.(Allotment); ok {
		rhsa := rhs.(Allotment)
		if len(lhsa) != len(rhsa) {
			return false
		}
		for i := range lhsa {
			if lhsa[i].Cmp(&rhsa[i]) != 0 {
				return false
			}
		}
	} else if lhsp, ok := lhs.(Portion); ok {
		rhsp := rhs.(Portion)
		return lhsp.Equals(&rhsp)
	} else if lhsf, ok := lhs.(Funding); ok {
		rhsf := rhs.(Funding)
		return lhsf.Equals(&rhsf)
	} else if lhs != rhs {
		return false
	}
	return true
}
