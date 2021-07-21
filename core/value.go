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
)

type Value interface {
	GetType() Type
}

type Account string

func (Account) GetType() Type { return TYPE_ACCOUNT }
func (a Account) String() string {
	return fmt.Sprintf("@%v", string(a))
}

type Asset string

func (Asset) GetType() Type { return TYPE_ASSET }
func (a Asset) String() string {
	return fmt.Sprintf("%v", string(a))
}

type Number uint64

func (Number) GetType() Type { return TYPE_NUMBER }
func (n Number) String() string {
	return fmt.Sprintf("%v", uint64(n))
}

type Monetary struct {
	Asset  Asset  `json:"asset"`
	Amount Amount `json:"amount"`
}

func (Monetary) GetType() Type { return TYPE_MONETARY }
func (a Monetary) String() string {
	return fmt.Sprintf("[%v %v]", a.Asset, a.Amount)
}

func (Allotment) GetType() Type { return TYPE_ALLOTMENT }

func (Amount) GetType() Type { return TYPE_AMOUNT }

func (Portion) GetType() Type { return TYPE_PORTION }

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
	} else if lhs != rhs {
		return false
	}
	return true
}
