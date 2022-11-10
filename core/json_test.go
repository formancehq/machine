package core

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccountTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "account",
		"value": "users:001"
	}`)
	value, err := NewValueFromTypedJSON(j)
	require.NoError(t, err)

	if !ValueEquals(*value, Account("users:001")) {
		t.Fatalf("unexpected value: %v", *value)
	}
}

func TestAssetTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "asset",
		"value": "EUR/2"
	}`)
	value, err := NewValueFromTypedJSON(j)
	require.NoError(t, err)

	if !ValueEquals(*value, Asset("EUR/2")) {
		t.Fatalf("unexpected value: %v", *value)
	}
}

func TestNumberTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "number",
		"value": 89849865111111111111111111111111111555555555555555555555555555555555555555555555555999999999999999999999
	}`)
	value, err := NewValueFromTypedJSON(j)
	require.NoError(t, err)

	num, err := ParseNumber("89849865111111111111111111111111111555555555555555555555555555555555555555555555555999999999999999999999")
	require.NoError(t, err)

	if !ValueEquals(*value, *num) {
		t.Fatalf("unexpected value: %v", *value)
	}
}

func TestMonetaryTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "monetary",
		"value": {
			"asset": "EUR/2",
			"amount": 123456
		}
	}`)
	value, err := NewValueFromTypedJSON(j)
	require.NoError(t, err)

	if !ValueEquals(*value, Monetary{
		Asset:  "EUR/2",
		Amount: *NewMonetaryInt(123456),
	}) {
		t.Fatalf("unexpected value: %v", *value)
	}
}

func TestPortionTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "portion",
		"value": "90%"
	}`)
	value, err := NewValueFromTypedJSON(j)
	require.NoError(t, err)

	portion, err := NewPortionSpecific(*big.NewRat(90, 100))
	require.NoError(t, err)

	if !ValueEquals(*value, *portion) {
		t.Fatalf("unexpected value: %v", *value)
	}
}

func TestInvalidTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"value": {
			"asset": "EUR/2",
			"amount": 123456
		}
	}`)
	_, err := NewValueFromTypedJSON(j)
	require.Error(t, err)
}
