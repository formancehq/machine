package core

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
	parts := allotment.Allocate(*NewMonetaryInt(15))
	expected_parts := []MonetaryInt{*NewMonetaryInt(13), *NewMonetaryInt(1), *NewMonetaryInt(1)}
	if len(parts) != len(expected_parts) {
		t.Fatalf("unexpected output %v != %v", parts, expected_parts)
	}
	for i := range parts {
		if !parts[i].Equal(&expected_parts[i]) {
			t.Fatalf("unexpected output %v != %v", parts, expected_parts)
		}
	}
}

func TestInvalidAllotments(t *testing.T) {
	_, err := NewAllotment([]Portion{
		{Remaining: true},
		{Specific: big.NewRat(2, 25)},
		{Remaining: true},
	})
	assert.Errorf(t, err, "allowed two remainings")

	_, err = NewAllotment([]Portion{
		{Specific: big.NewRat(1, 2)},
		{Specific: big.NewRat(1, 2)},
		{Specific: big.NewRat(1, 2)},
	})
	assert.Errorf(t, err, "allowed more than 100%")

	_, err = NewAllotment([]Portion{
		{Specific: big.NewRat(1, 2)},
		{Specific: big.NewRat(1, 2)},
		{Remaining: true},
	})
	assert.Errorf(t, err, "allowed remaining but already at 100%")
}
