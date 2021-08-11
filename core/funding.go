package core

import (
	"errors"
	"fmt"
)

type FundingPart struct {
	Amount  uint64
	Account Account
}

type Funding struct {
	Asset    Asset
	Parts    []FundingPart
	Infinite bool
}

func (lhs *Funding) Equals(rhs *Funding) bool {
	if lhs.Asset != rhs.Asset {
		return false
	}
	if len(lhs.Parts) != len(rhs.Parts) {
		return false
	}
	for i := range lhs.Parts {
		if lhs.Parts[i] != rhs.Parts[i] {
			return false
		}
	}
	return true
}

func (f Funding) String() string {
	out := fmt.Sprintf("[%v", string(f.Asset))
	for _, part := range f.Parts {
		out += fmt.Sprintf(" %v %v", part.Amount, part.Account)
	}
	return out + "]"
}

func (f Funding) Take(amount uint64) (Funding, Funding, error) {
	result := Funding{
		Asset: f.Asset,
	}
	remainder := Funding{
		Asset: f.Asset,
	}
	remaining_to_withdraw := amount
	i := 0
	for remaining_to_withdraw > 0 && i < len(f.Parts) {
		amt_to_withdraw := f.Parts[i].Amount
		// if this part has excess balance, put it in the remainder & only take what's needed
		if amt_to_withdraw > remaining_to_withdraw {
			rem := amt_to_withdraw - remaining_to_withdraw
			amt_to_withdraw = remaining_to_withdraw
			remainder.Parts = append(remainder.Parts, FundingPart{
				Account: f.Parts[i].Account,
				Amount:  rem,
			})
		}
		remaining_to_withdraw -= amt_to_withdraw
		result.Parts = append(result.Parts, FundingPart{
			Account: f.Parts[i].Account,
			Amount:  amt_to_withdraw,
		})
		i++
	}
	for i < len(f.Parts) {
		remainder.Parts = append(remainder.Parts, FundingPart{
			Account: f.Parts[i].Account,
			Amount:  f.Parts[i].Amount,
		})
		i++
	}
	if remaining_to_withdraw != 0 {
		return Funding{}, Funding{}, errors.New("insufficient funding")
	}
	return result, remainder, nil
}

func (f Funding) Total() uint64 {
	total := uint64(0)
	for _, part := range f.Parts {
		total += part.Amount
	}
	return total
}
