package core

import (
	"errors"
	"fmt"
	"math/big"
)

type Allotment []big.Rat

func (Allotment) GetType() Type { return TYPE_ALLOTMENT }

func NewAllotment(portions []Portion) (*Allotment, error) {
	n := len(portions)
	total := big.NewRat(0, 1)
	var remaining_idx *int
	allotment := make([]big.Rat, n)
	for i := 0; i < n; i++ {
		if portions[i].Remaining {
			if remaining_idx != nil {
				return nil, errors.New("two uses of `remaining` in the same allotment")
			}
			allotment[i] = big.Rat{} // temporary
			idx := i
			remaining_idx = &idx
		} else {
			rat := big.Rat(*portions[i].Specific)
			allotment[i] = rat
			total.Add(total, &rat)
		}
	}
	if total.Cmp(big.NewRat(1, 1)) == 1 {
		return nil, errors.New("sum of portions exceeded 100%")
	}
	if remaining_idx != nil && total.Cmp(big.NewRat(1, 1)) != -1 {
		return nil, errors.New("portions include 'remaining' but sum is 100% or more")
	}
	if remaining_idx != nil {
		remaining := big.NewRat(1, 1)
		remaining.Sub(remaining, total)
		allotment[*remaining_idx] = *remaining
	}
	result := Allotment(allotment)
	return &result, nil
}

func (a Allotment) String() string {
	out := "{ "
	for i, ratio := range a {
		out += fmt.Sprintf("%v", &ratio)
		if i != len(a)-1 {
			out += " : "
		}
	}
	return out + " }"
}

func (a Allotment) Allocate(amount MonetaryInt) []MonetaryInt {
	amt_bigint := big.Int(amount)
	parts := make([]MonetaryInt, len(a))
	total_allocated := *NewMonetaryInt(0)
	// for every part in the allotment, calculate the floored value
	for i, allot := range a {
		var res big.Int
		res.Mul(&amt_bigint, allot.Num())
		res.Div(&res, allot.Denom())
		parts[i] = MonetaryInt(res)
		total_allocated = *total_allocated.Add(&parts[i])
	}
	for i := range parts {
		if total_allocated.Lt(&amount) {
			parts[i] = *parts[i].Add(NewMonetaryInt(1))
			total_allocated = *total_allocated.Add(NewMonetaryInt(1))
		}
	}
	return parts
}
