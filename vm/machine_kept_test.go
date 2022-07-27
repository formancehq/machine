package vm

import (
	"testing"

	ledger "github.com/numary/ledger/pkg/core"
	"github.com/numary/machine/core"
)

func TestKeptDestinationAllotment(t *testing.T) {
	testJSON(t,
		`
send [GEM 100] (
	source = {
		@a
		@world
	}
	destination = {
		50% kept
		25% to @x
		25% to @y
	}
)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"a": {
				"GEM": 1,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      1,
					Source:      "a",
					Destination: "x",
				},
				{
					Asset:       "GEM",
					Amount:      24,
					Source:      "world",
					Destination: "x",
				},
				{
					Asset:       "GEM",
					Amount:      25,
					Source:      "world",
					Destination: "y",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}

func TestKeptComplex(t *testing.T) {
	testJSON(t,
		`
		send [GEM 100] (
			source = {
				@foo
				@bar
				@baz
			}
			destination = {
				50% to {
					max [GEM 8] to {
						50% kept
						25% to @arst
						25% kept
					}
					remaining to @thing
				}
				20% to @qux
				5% kept
				remaining to @quz
			}
		)`,
		`{}`,
		map[string]map[string]core.Value{},
		map[string]map[string]uint64{
			"foo": {
				"GEM": 20,
			},
			"bar": {
				"GEM": 40,
			},
			"baz": {
				"GEM": 40,
			},
		},
		CaseResult{
			Printed: []core.Value{},
			Postings: []ledger.Posting{
				{
					Asset:       "GEM",
					Amount:      2,
					Source:      "foo",
					Destination: "arst",
				},
				{
					Asset:       "GEM",
					Amount:      18,
					Source:      "foo",
					Destination: "thing",
				},
				{
					Asset:       "GEM",
					Amount:      24,
					Source:      "bar",
					Destination: "thing",
				},
				{
					Asset:       "GEM",
					Amount:      16,
					Source:      "bar",
					Destination: "qux",
				},
				{
					Asset:       "GEM",
					Amount:      4,
					Source:      "baz",
					Destination: "qux",
				},
				{
					Asset:       "GEM",
					Amount:      25,
					Source:      "baz",
					Destination: "quz",
				},
			},
			ExitCode: EXIT_OK,
		},
	)
}
