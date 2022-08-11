package core

import (
	"errors"
	"math/big"
)

type Number big.Int

func (Number) isValue()      {}
func (Number) GetType() Type { return TYPE_NUMBER }

func (a *Number) Add(b *Number) *Number {
	if a == nil {
		a = NewNumber(0)
	}

	if b == nil {
		b = NewNumber(0)
	}

	return (*Number)(big.NewInt(0).Add((*big.Int)(a), (*big.Int)(b)))
}

func (a *Number) Sub(b *Number) *Number {
	if a == nil {
		a = NewNumber(0)
	}

	if b == nil {
		b = NewNumber(0)
	}

	return (*Number)(big.NewInt(0).Sub((*big.Int)(a), (*big.Int)(b)))
}

func (a *Number) Neg() *Number {
	return (*Number)(big.NewInt(0).Neg((*big.Int)(a)))
}

func (a *Number) OrZero() *Number {
	if a == nil {
		return NewNumber(0)
	}

	return a
}

func (a *Number) Lte(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) <= 0
}

func (a *Number) Gte(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) >= 0
}

func (a *Number) Lt(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) < 0
}

func (a *Number) Ltz() bool {
	return (*big.Int)(a).Cmp(big.NewInt(0)) < 0
}

func (a *Number) Gt(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) > 0
}

func (a *Number) Eq(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) == 0
}

func (a *Number) Equal(b *Number) bool {
	return (*big.Int)(a).Cmp((*big.Int)(b)) == 0
}

func (a *Number) Cmp(b *Number) int {
	return (*big.Int)(a).Cmp((*big.Int)(b))
}

func (a *Number) Uint64() uint64 {
	return (*big.Int)(a).Uint64()
}

func (a Number) String() string {
	return (*big.Int)(&a).String()
}

func (a *Number) UnmarshalJSON(b []byte) error {
	return (*big.Int)(a).UnmarshalJSON(b)
}

func (a *Number) MarshalJSON() ([]byte, error) {
	if a == nil {
		return []byte("0"), nil
	}
	return (*big.Int)(a).MarshalJSON()
}

func (a *Number) MarshalText() ([]byte, error) {
	return (*big.Int)(a).MarshalText()
}

func (a *Number) UnmarshalText(b []byte) error {
	return (*big.Int)(a).UnmarshalText(b)
}

func NewNumber(i int64) *Number {
	return (*Number)(big.NewInt(i))
}

func ParseNumber(s string) (*Number, error) {
	i, ok := big.NewInt(0).SetString(s, 10)

	if !ok {
		return nil, errors.New("invalid monetary int")
	}

	return (*Number)(i), nil
}
