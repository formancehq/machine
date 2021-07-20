package core

import (
	"fmt"
	"math/big"
	"reflect"
)

type Type byte

const (
	TYPE_ACCOUNT = Type(iota + 1)
	TYPE_ASSET
	TYPE_NUMBER
	TYPE_MONETARY
	TYPE_ALLOTMENT
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
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

func (Monetary) GetType() Type { return TYPE_MONETARY }
func (a Monetary) String() string {
	return fmt.Sprintf("[%v %v]", a.Asset, a.Amount)
}

type Allotment []big.Rat

func (Allotment) GetType() Type { return TYPE_ALLOTMENT }
func (a Allotment) String() string {
	out := "{\n"
	for _, ratio := range a {
		out += fmt.Sprintf("	%v\n", &ratio)
	}
	return out + "}"
}

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
