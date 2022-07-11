package compiler

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

type CaseResult struct {
	Instructions []byte
	Resources    []program.Resource
	Variables    []string
	Error        string
}

type TestCase struct {
	Case     string
	Expected CaseResult
}

func test(t *testing.T, c TestCase) {
	p, err := Compile(c.Case)

	if c.Expected.Error != "" {
		if err == nil {
			t.Error(errors.New("expected error and got none"))
			return
		} else if err.Error() == "" {
			t.Error(errors.New("error was not fed to the error listener"))
		} else if !strings.Contains(err.Error(), c.Expected.Error) {
			t.Error(fmt.Errorf("error is not the one expected: %v", err))
			return
		}
	} else {
		if err != nil {
			t.Error(fmt.Errorf("did not expect error: %v", err))
			return
		} else if p == nil {
			t.Error(errors.New("did not receive any output"))
			return
		} else if !bytes.Equal(p.Instructions, c.Expected.Instructions) {
			t.Error(fmt.Errorf("generated program is incorrect: %v", *p))
			fmt.Println("has:", p.Instructions, "want:", c.Expected.Instructions)
			return
		} else if len(p.Resources) != len(c.Expected.Resources) {
			t.Error(fmt.Errorf("unexpected program resources (=/= lengths): %v", *p))
			return
		} else {
			for i := range c.Expected.Resources {
				if !checkResourceEquals(t, p.Resources[i], c.Expected.Resources[i]) {
					t.Error(fmt.Errorf("%v: %v is not %v: %v",
						p.Resources[i], reflect.TypeOf(p.Resources[i]).Name(),
						c.Expected.Resources[i], reflect.TypeOf(c.Expected.Resources[i]).Name(),
					))
					t.Error(fmt.Errorf("unexpected program resources: %v", *p))
					return
				}
			}
		}
	}
}

func checkResourceEquals(t *testing.T, res program.Resource, expected program.Resource) bool {
	if reflect.TypeOf(res) != reflect.TypeOf(expected) {
		return false
	}
	switch res := res.(type) {
	case program.Constant:
		if reflect.TypeOf(res.Inner).Kind() == reflect.Ptr {
			t.Fatal("generated program contained a constant with a pointer value")
		}
		return core.ValueEquals(res.Inner, expected.(program.Constant).Inner)
	case program.Parameter:
		e := expected.(program.Parameter)
		return res.Typ == e.Typ && res.Name == e.Name
	case program.Metadata:
		e := expected.(program.Metadata)
		return res.SourceAccount == e.SourceAccount &&
			res.Key == e.Key &&
			res.Typ == e.Typ
	default:
		panic("invalid resource")
	}
}

func TestSimplePrint(t *testing.T) {
	test(t, TestCase{
		Case: "print 1",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_PRINT,
			},
			Resources: []program.Resource{},
			Error:     "",
		},
	})
}

func TestCompositeExpr(t *testing.T) {
	test(t, TestCase{
		Case: "print 29 + 15 - 2",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_IPUSH, 29, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IPUSH, 15, 00, 00, 00, 00, 00, 00, 00,
				program.OP_IADD,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_ISUB,
				program.OP_PRINT,
			},
			Resources: []program.Resource{},
			Error:     "",
		},
	})
}

func TestFail(t *testing.T) {
	test(t, TestCase{
		Case: "fail",
		Expected: CaseResult{
			Instructions: []byte{program.OP_FAIL},
			Resources:    []program.Resource{},
			Error:        "",
		},
	})
}

func TestCRLF(t *testing.T) {
	test(t, TestCase{
		Case: "print @a\r\nprint @b",
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00, program.OP_PRINT,
				program.OP_APUSH, 01, 00, program.OP_PRINT,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Account("a")},
				program.Constant{Inner: core.Account("b")},
			},
			Error: "",
		},
	})
}

func TestConstant(t *testing.T) {
	user := core.Account("user:U001")
	test(t, TestCase{
		Case: "print @user:U001",
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{program.Constant{Inner: user}},
			Error:        "",
		},
	})
}

func TestSetTxMeta(t *testing.T) {
	test(t, TestCase{
		Case: `
		set_tx_meta("beneficiary", @platform)
		`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_TX_META,
			},
			Resources: []program.Resource{
				program.Constant{
					Inner: core.Account("platform"),
				},
				program.Constant{
					Inner: core.String("beneficiary"),
				},
			},
			Error: "",
		},
	})
}

func TestSetTxMetaVars(t *testing.T) {
	test(t, TestCase{
		Case: `
		vars {
			portion $commission
		}
		set_tx_meta("fee", $commission)
		`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_TX_META,
			},
			Resources: []program.Resource{
				program.Parameter{
					Typ:  core.TYPE_PORTION,
					Name: "commission",
				},
				program.Constant{
					Inner: core.String("fee"),
				},
			},
			Error: "",
		},
	})
}

func TestComments(t *testing.T) {
	test(t, TestCase{
		Case: `
		/* This is a multi-line comment, it spans multiple lines
		and /* doesn't choke on nested comments */ ! */
		vars {
			account $a
		}
		// this is a single-line comment
		print $a
		`,
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{program.Parameter{Typ: core.TYPE_ACCOUNT, Name: "a"}},
			Error:        "",
		},
	})
}

func TestUndeclaredVariable(t *testing.T) {
	test(t, TestCase{
		Case: "print $nope",
		Expected: CaseResult{
			Instructions: []byte{},
			Resources:    []program.Resource{},
			Error:        "declared",
		},
	})
}

func TestInvalidTypeInSendValue(t *testing.T) {
	test(t, TestCase{
		Case: `
		send @a (
			source = {
				@a
				[GEM 2]
			}
			destination = @b
		)`,
		Expected: CaseResult{
			Instructions: []byte{},
			Resources:    []program.Resource{},
			Error:        "wrong type",
		},
	})
}

func TestInvalidTypeInSource(t *testing.T) {
	test(t, TestCase{
		Case: `
		send [USD/2 99] (
			source = {
				@a
				[GEM 2]
			}
			destination = @b
		)`,
		Expected: CaseResult{
			Instructions: []byte{},
			Resources:    []program.Resource{},
			Error:        "wrong type",
		},
	})
}

func TestAllocationFractions(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 43] (
			source = @foo
			destination = {
				1/8 to @bar
				7/8 to @baz
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_BUMP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: 43,
				}},
				program.Constant{Inner: core.Account("foo")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(7, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
				program.Constant{Inner: core.Account("bar")},
				program.Constant{Inner: core.Account("baz")},
			},
			Error: "",
		},
	})
}

func TestComplexDestination(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 43] (
			source = @a
			destination = {
				1/8 to {
					max [EUR/2 10] to @b
					@c
				}
				7/8 to @d
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_BUMP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 04, 00,
				program.OP_TAKE_MAX,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
				program.OP_APUSH, 06, 00,
				program.OP_SEND,
				program.OP_APUSH, 07, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: 43,
				}},
				program.Constant{Inner: core.Account("a")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(7, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: 10,
				}},
				program.Constant{Inner: core.Account("b")},
				program.Constant{Inner: core.Account("c")},
				program.Constant{Inner: core.Account("d")},
			},
			Error: "",
		},
	})
}

func TestAllocationPercentages(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 43] (
			source = @foo
			destination = {
				12.5% to @bar
				37.5% to @baz
				50% to @qux
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_BUMP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 03, 00,
				program.OP_APUSH, 04, 00,
				program.OP_IPUSH, 03, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
				program.OP_APUSH, 06, 00,
				program.OP_SEND,
				program.OP_APUSH, 07, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: 43,
				}},
				program.Constant{Inner: core.Account("foo")},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 2)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(3, 8)}},
				program.Constant{Inner: core.Portion{Specific: big.NewRat(1, 8)}},
				program.Constant{Inner: core.Account("bar")},
				program.Constant{Inner: core.Account("baz")},
				program.Constant{Inner: core.Account("qux")},
			},
			Error: "",
		},
	})
}

func TestSend(t *testing.T) {
	alice := core.Account("alice")
	bob := core.Account("bob")
	test(t, TestCase{
		Case: `send [EUR/2 99] (
	source = @alice
	destination = @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_BUMP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "EUR/2", Amount: 99}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob}},
			Error: "",
		},
	})
}

func TestSendAll(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 *] (
	source = @alice
	destination = @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 01, 00,
				program.OP_APUSH, 00, 00,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Resources: []program.Resource{
				program.Constant{Inner: core.Asset("EUR/2")},
				program.Constant{Inner: core.Account("alice")},
				program.Constant{Inner: core.Account("bob")}},
			Error: "",
		},
	})
}

func TestMetadata(t *testing.T) {
	test(t, TestCase{
		Case: `
		vars {
			account $sale
			account $seller = meta($sale, "seller")
			portion $commission = meta($seller, "commission")
		}
		send [EUR/2 53] (
			source = $sale
			destination = {
				remaining to $seller
				$commission to @platform
			}
		)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 03, 00,
				program.OP_ASSET,
				program.OP_TAKE_ALL,
				program.OP_APUSH, 03, 00,
				program.OP_TAKE,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_BUMP,
				program.OP_REPAY,
				program.OP_APUSH, 02, 00,
				program.OP_APUSH, 04, 00,
				program.OP_IPUSH, 02, 00, 00, 00, 00, 00, 00, 00,
				program.OP_MAKE_ALLOTMENT,
				program.OP_ALLOC,
				program.OP_APUSH, 01, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			}, Resources: []program.Resource{
				program.Parameter{Typ: core.TYPE_ACCOUNT, Name: "sale"},
				program.Metadata{Typ: core.TYPE_ACCOUNT, SourceAccount: core.NewAddress(0), Key: "seller"},
				program.Metadata{Typ: core.TYPE_PORTION, SourceAccount: core.NewAddress(1), Key: "commission"},
				program.Constant{Inner: core.Monetary{Asset: "EUR/2", Amount: 53}},
				program.Constant{Inner: core.NewPortionRemaining()},
				program.Constant{Inner: core.Account("platform")},
			},
			Error: "",
		},
	})
}

func TestSyntaxError(t *testing.T) {
	test(t, TestCase{
		Case: "print fail",
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "mismatched input",
		},
	})
}

func TestLogicError(t *testing.T) {
	test(t, TestCase{
		Case: `send [EUR/2 200] (
			source = 200
			destination = @bob
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "expected",
		},
	})
}

func TestPreventTakeAllFromWorld(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM *] (
			source = @world
			destination = @foo
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "cannot",
		},
	})
}

func TestPreventAddToBottomlessSource(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM 1000] (
			source = {
				@a
				@world
				@c
			}
			destination = @out
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "world",
		},
	})
}

func TestPreventAddToBottomlessSource2(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM 1000] (
			source = {
				{
					@a
					@world
				}
				{
					@b
					@world
				}
			}
			destination = @out
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "world",
		},
	})
}

func TestPreventSourceAlreadyEmptied(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM 1000] (
			source = {
				{
					@a
					@world
				}
				@a
			}
			destination = @out
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "@a",
		},
	})
}

func TestPreventTakeAllFromAllocation(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM *] (
			source = {
				50% from @a
				50% from @b
			}
			destination = @out
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "all",
		},
	})
}

func TestOverflowingAllocation(t *testing.T) {
	fmt.Println("case: >100%")
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = {
				2/3 to @a
				2/3 to @b
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "100%",
		},
	})

	fmt.Println("case: =100% + remaining")
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = {
				1/2 to @a
				1/2 to @b
				remaining to @c
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "100%",
		},
	})

	fmt.Println("case: >100% + remaining")
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = {
				2/3 to @a
				1/2 to @b
				remaining to @c
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "100%",
		},
	})

	fmt.Println("case: const remaining + remaining")
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = {
				2/3 to @a
				remaining to @b
				remaining to @c
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "`remaining` in the same",
		},
	})

	fmt.Println("case: dyn remaining + remaining")
	test(t, TestCase{
		Case: `
		vars {
			portion $p
		}
		send [GEM 15] (
			source = @world
			destination = {
				$p to @a
				remaining to @b
				remaining to @c
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "`remaining` in the same",
		},
	})

	fmt.Println("case: >100% + remaining + variable")
	test(t, TestCase{
		Case: `
		vars {
			portion $prop
		}
		send [GEM 15] (
			source = @world
			destination = {
				1/2 to @a
				2/3 to @b
				remaining to @c
				$prop to @d
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "100%",
		},
	})

	fmt.Println("case: variable - remaining")
	test(t, TestCase{
		Case: `
		vars {
			portion $prop
		}
		send [GEM 15] (
			source = @world
			destination = {
				2/3 to @a
				$prop to @b
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "remaining",
		},
	})
}

func TestAllocationWrongDestination(t *testing.T) {
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = [GEM 10]
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "account",
		},
	})
	test(t, TestCase{
		Case: `send [GEM 15] (
			source = @world
			destination = {
				2/3 to @a
				1/3 to [GEM 10]
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "account",
		},
	})
}

func TestAllocationInvalidPortion(t *testing.T) {
	test(t, TestCase{
		Case: `
		vars {
			account $p
		}
		send [GEM 15] (
			source = @world
			destination = {
				10% to @a
				$p to @b
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "type",
		},
	})
}

func TestCappedDestination(t *testing.T) {
	test(t, TestCase{
		Case: `
		vars {
			account $p
		}
		send [GEM 15] (
			source = @world
			destination = {
				50% to max [GEM 15] to @a
				50% to @b
			}
		)`,
		Expected: CaseResult{
			Instructions: nil,
			Resources:    nil,
			Error:        "cap",
		},
	})
}

// func TestTooManyConstants(t *testing.T) {
// 	script := ""
// 	for i := 0; i < 11000; i++ {
// 		script += fmt.Sprintf(`
// 		send [A%d 0] (
// 			source=@a%d
// 			destination=@b%d
// 		)`, i, i, i)
// 		script += "\n"
// 	}
// 	test(t, TestCase{
// 		Case: script,
// 		Expected: CaseResult{
// 			Instructions: nil,
// 			Constants:    nil,
// 			Error:        "exceeded",
// 		},
// 	})
// }
