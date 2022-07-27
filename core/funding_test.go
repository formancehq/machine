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

func TestFundingTakeInfinite1(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  70,
			},
		},
		Infinite: true,
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
				Account: Account("world"),
				Amount:  10,
			},
		},
	}
	if !ValueEquals(result, expected_result) {
		t.Fatalf("unexpected result: %v", result)
	}
	expected_remainder := Funding{
		Asset:    Asset("COIN"),
		Infinite: true,
	}
	if !ValueEquals(remainder, expected_remainder) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}

func TestFundingTakeInfinite2(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  70,
			},
		},
		Infinite: true,
	}
	result, remainder, err := f.Take(30)
	if err != nil {
		t.Fatal(err)
	}
	expected_result := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  30,
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
				Account: Account("aaa"),
				Amount:  40,
			},
		},
		Infinite: true,
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
				Amount:  30,
			},
		},
	}
	result, remainder := f.TakeMax(80)
	if !ValueEquals(result, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  30,
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
				Amount:  90,
			},
		},
	}
	result, remainder := f.TakeMax(80)
	if !ValueEquals(result, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  80,
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
				Amount:  10,
			},
		},
	}) {
		t.Fatalf("unexpected remainder: %v", remainder)
	}
}

func TestFundingTakeMaxInfinite(t *testing.T) {
	f := Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  20,
			},
		},
		Infinite: true,
	}
	result, remainder := f.TakeMax(80)
	if !ValueEquals(result, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("aaa"),
				Amount:  20,
			},
			{
				Account: Account("world"),
				Amount:  60,
			},
		},
	}) {
		t.Fatalf("unexpected result: %v", result)
	}
	if !ValueEquals(remainder, Funding{
		Asset:    Asset("COIN"),
		Parts:    []FundingPart{},
		Infinite: true,
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
				Amount:  10,
			},
			{
				Account: Account("bbb"),
				Amount:  20,
			},
			{
				Account: Account("ccc"),
				Amount:  30,
			},
		},
	}
	rev, err := f.Reverse()
	if err != nil {
		t.Fatal(err)
	}
	if !ValueEquals(*rev, Funding{
		Asset: Asset("COIN"),
		Parts: []FundingPart{
			{
				Account: Account("ccc"),
				Amount:  30,
			},
			{
				Account: Account("bbb"),
				Amount:  20,
			},
			{
				Account: Account("aaa"),
				Amount:  10,
			},
		},
	}) {
		t.Fatalf("unexpected result: %v", rev)
	}
}

func TestInfiniteFundingReversal(t *testing.T) {
	_, err := Funding{
		Asset:    Asset("COIN"),
		Parts:    []FundingPart{},
		Infinite: true,
	}.Reverse()
	if err == nil {
		t.Fatal("expected error")
	}
}
