package core

import (
	"testing"
)

func TestFundingTake(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  70,
			},
			{
				Account: Account("bbb"),
				Amount:  30,
			},
			{
				Account: Account("ccc"),
				Amount:  50,
			},
		},
	}
	result, remainder, err := f.Take(80)
	if err != nil {
		t.Fatal(err)
	}
	expected_result := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  70,
			},
			{
				Account: Account("bbb"),
				Amount:  10,
			},
		},
	}
	if !ValueEquals(result, expected_result) {
		t.Fatalf("unexpected result: %v", result)
	}
	expected_remainder := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("bbb"),
				Amount:  20,
			},
			{
				Account: Account("ccc"),
				Amount:  50,
			},
		},
	}
	if !ValueEquals(remainder, expected_remainder) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}
