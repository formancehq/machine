package core

import (
	"fmt"
	"math/big"
)

type Type byte

const (
	TYPE_ACCOUNT = Type(iota + 1)
	TYPE_ASSET
	TYPE_NUMBER
	TYPE_MONETARY
	TYPE_ALLOTMENT
	TYPE_AMOUNT
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

type Allotment []big.Rat

func (Allotment) GetType() Type { return TYPE_ALLOTMENT }
func (a Allotment) String() string {
	out := "{\n"
	for _, ratio := range a {
		out += fmt.Sprintf("	%v\n", &ratio)
	}
	return out + "}"
}

// invariant: specific must be zero if all is true
type Amount struct {
	All      bool
	Specific uint64
}

func (Amount) GetType() Type { return TYPE_AMOUNT }
func (a Amount) String() string {
	if a.All {
		return "*"
	} else {
		return fmt.Sprint(a.Specific)
	}
}

func NewAmountAll() Amount {
	return Amount{
		All:      true,
		Specific: 0,
	}
}

func NewAmountSpecific(x uint64) Amount {
	return Amount{
		All:      false,
		Specific: x,
	}
}
