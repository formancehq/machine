package core

import (
	"math/big"
	"testing"
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
	parts := allotment.Allocate(15)
	expected_parts := []uint64{13, 1, 1}
	if len(parts) != len(expected_parts) {
		t.Fatalf("unexpected output %v != %v", parts, expected_parts)
	}
	for i := range parts {
		if parts[i] != expected_parts[i] {
			t.Fatalf("unexpected output %v != %v", parts, expected_parts)
		}
	}
}
