package core

import (
	"errors"
	"fmt"
	"math/big"
)

type Allotment []big.Rat

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
	out := "{\n"
	for _, ratio := range a {
		out += fmt.Sprintf("	%v\n", &ratio)
	}
	return out + "}"
}
