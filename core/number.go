package core

import (
	"math/big"
)

type Number = MonetaryInt

func NewNumber(i int64) *Number {
	return (*Number)(big.NewInt(i))
}

func ParseNumber(s string) (*Number, error) {
	return ParseMonetaryInt(s)
}
