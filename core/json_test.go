package core

import (
	"encoding/json"
	"testing"

	ledger "github.com/numary/ledger/pkg/core"
)

func TestAccountTypedJSON(t *testing.T) {
	j := json.RawMessage(`{
		"type": "account",
		"value": "users:001"
	}`)
	value, err := NewValueFromTypedJSON(j)
	if err != nil {
		t.Fatal(err)
	}
	if !ValueEquals(*value, Account("users:001")) {
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
	if err != nil {
		t.Fatal(err)
	}
	if !ValueEquals(*value, Monetary{
		Asset:  Asset("EUR/2"),
		Amount: *ledger.NewMonetaryInt(123456),
	}) {
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
	if err == nil {
		t.Fatalf("error expected but got none")
	}
}
