package core

import "fmt"

type Amount struct {
	All      bool
	Specific uint64 // invariant: Specific must be zero if All is true
}

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
