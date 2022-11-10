package core

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
)

type Portion struct {
	Remaining bool     `json:"remaining"`
	Specific  *big.Rat `json:"specific"`
}

func (Portion) GetType() Type { return TypePortion }

func NewPortionRemaining() Portion {
	return Portion{
		Remaining: true,
		Specific:  nil,
	}
}

func NewPortionSpecific(r big.Rat) (*Portion, error) {
	if r.Cmp(big.NewRat(0, 1)) != 1 || r.Cmp(big.NewRat(1, 1)) != -1 {
		return nil, errors.New("portion must be between 0% and 100% exclusive")
	}
	return &Portion{
		Remaining: false,
		Specific:  &r,
	}, nil
}

func (lhs Portion) Equals(rhs Portion) bool {
	if lhs.Remaining != rhs.Remaining {
		return false
	}
	if !lhs.Remaining && lhs.Specific.Cmp(rhs.Specific) != 0 {
		return false
	}
	return true
}

func ParsePortionSpecific(input string) (*Portion, error) {
	var res *big.Rat

	re := regexp.MustCompile(`^([0-9]+)(?:[.]([0-9]+))?[%]$`)
	percentMatch := re.FindStringSubmatch(input)
	if len(percentMatch) != 0 {
		integral := percentMatch[1]
		fractional := percentMatch[2]
		rat, ok := new(big.Rat).SetString(integral + "." + fractional)
		if !ok {
			return nil, errors.New("invalid format")
		}
		rat.Mul(rat, big.NewRat(1, 100))
		res = rat
	} else {
		re = regexp.MustCompile(`^([0-9]+)\s?[/]\s?([0-9]+)$`)
		fractionMatch := re.FindStringSubmatch(input)
		if len(fractionMatch) != 0 {
			numerator := fractionMatch[1]
			denominator := fractionMatch[2]
			rat, ok := new(big.Rat).SetString(numerator + "/" + denominator)
			if !ok {
				return nil, errors.New("invalid format")
			}
			res = rat
		}
	}

	if res == nil {
		return nil, errors.New("invalid format")
	}
	portion, err := NewPortionSpecific(*res)
	if err != nil {
		return nil, err
	}
	return portion, nil
}

func (p Portion) String() string {
	if p.Remaining {
		return "remaining"
	} else {
		return fmt.Sprintf("%v", p.Specific)
	}
}
