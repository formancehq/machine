package core

import (
	"testing"

	ledger "github.com/numary/ledger/pkg/core"
)

func TestFundingTake(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(70),
			},
			{
				Account: Account("bbb"),
				Amount:  *ledger.NewMonetaryInt(30),
			},
			{
				Account: Account("ccc"),
				Amount:  *ledger.NewMonetaryInt(50),
			},
		},
	}
	result, remainder, err := f.Take(*ledger.NewMonetaryInt(80))
	if err != nil {
		t.Fatal(err)
	}
	expected_result := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(70),
			},
			{
				Account: Account("bbb"),
				Amount:  *ledger.NewMonetaryInt(10),
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
				Amount:  *ledger.NewMonetaryInt(20),
			},
			{
				Account: Account("ccc"),
				Amount:  *ledger.NewMonetaryInt(50),
			},
		},
	}
	if !ValueEquals(remainder, expected_remainder) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}

func TestFundingTakeMaxUnder(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(30),
			},
		},
	}
	result, remainder := f.TakeMax(*ledger.NewMonetaryInt(80))
	if !ValueEquals(result, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(30),
			},
		},
	}) {
		t.Fatalf("unexpected result: %v", result)
	}
	if !ValueEquals(remainder, Funding{
		Asset: Asset("COIN"),
	}) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}

func TestFundingTakeMaxAbove(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(90),
			},
		},
	}
	result, remainder := f.TakeMax(*ledger.NewMonetaryInt(80))
	if !ValueEquals(result, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(80),
			},
		},
	}) {
		t.Fatalf("unexpected result: %v", result)
	}
	if !ValueEquals(remainder, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(10),
			},
		},
	}) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}

func TestFundingReversal(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(10),
			},
			{
				Account: Account("bbb"),
				Amount:  *ledger.NewMonetaryInt(20),
			},
			{
				Account: Account("ccc"),
				Amount:  *ledger.NewMonetaryInt(30),
			},
		},
	}
	rev := f.Reverse()
	if !ValueEquals(rev, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("ccc"),
				Amount:  *ledger.NewMonetaryInt(30),
			},
			{
				Account: Account("bbb"),
				Amount:  *ledger.NewMonetaryInt(20),
			},
			{
				Account: Account("aaa"),
				Amount:  *ledger.NewMonetaryInt(10),
			},
		},
	}) {
		t.Fatalf("unexpected result: %v", rev)
	}
}
