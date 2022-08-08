package core

import (
	"math/big"
	"testing"

	ledger "github.com/numary/ledger/pkg/core"
)

func TestAllocate(t *testing.T) {
	allotment, err := NewAllotment([]Portion{
		{Specific: big.NewRat(4, 5)},
		{Specific: big.NewRat(2, 25)},
		{Specific: big.NewRat(3, 25)},
	})
	if err != nil {
		t.Fatal(err)
	}
	parts := allotment.Allocate(*ledger.NewMonetaryInt(15))
	expected_parts := []ledger.MonetaryInt{*ledger.NewMonetaryInt(13), *ledger.NewMonetaryInt(1), *ledger.NewMonetaryInt(1)}
	if len(parts) != len(expected_parts) {
		t.Fatalf("unexpected output %v != %v", parts, expected_parts)
	}
	for i := range parts {
		if !parts[i].Equal(&expected_parts[i]) {
			t.Fatalf("unexpected output %v != %v", parts, expected_parts)
		}
	}
}
