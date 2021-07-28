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
			fmt.Println(p.Instructions, "vs", c.Expected.Instructions)
			return
		} else if len(p.Resources) != len(c.Expected.Resources) {
			t.Error(fmt.Errorf("unexpected program constants: %v", *p))
			return
		} else {
			for i := range c.Expected.Resources {
				if !ResourceEquals(p.Resources[i], c.Expected.Resources[i]) {
					t.Error(fmt.Errorf("unexpected program constants: %v", *p))
					return
				}
			}
		}
	}
}

func ResourceEquals(lhs program.Resource, rhs program.Resource) bool {
	if reflect.TypeOf(lhs) != reflect.TypeOf(rhs) {
		return false
	}
	switch lhs := lhs.(type) {
	case program.Constant:
		return core.ValueEquals(lhs.Inner, rhs.(program.Constant).Inner)
	case program.Parameter:
		return lhs.Typ == rhs.(program.Parameter).Typ
	case program.Metadata:
		return lhs.Typ == rhs.(program.Metadata).Typ
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

func TestConstant(t *testing.T) {
	user := core.Account("user:001")
	test(t, TestCase{
		Case: "print @user:001",
		Expected: CaseResult{
			Instructions: []byte{program.OP_APUSH, 00, 00, program.OP_PRINT},
			Resources:    []program.Resource{program.Constant{Inner: user}},
			Error:        "",
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
			Resources:    []program.Resource{program.Parameter{Typ: core.TYPE_ACCOUNT}},
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
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: core.NewAmountSpecific(43),
				}},
				program.Constant{Inner: core.Account("foo")},
				program.Constant{Inner: core.Allotment{*big.NewRat(1, 8), *big.NewRat(7, 8)}},
				program.Constant{Inner: core.Account("bar")},
				program.Constant{Inner: core.Account("baz")},
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
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_ALLOC,
				program.OP_APUSH, 03, 00,
				program.OP_SEND,
				program.OP_APUSH, 04, 00,
				program.OP_SEND,
				program.OP_APUSH, 05, 00,
				program.OP_SEND,
			},
			Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{
					Asset:  "EUR/2",
					Amount: core.NewAmountSpecific(43),
				}},
				program.Constant{Inner: core.Account("foo")},
				program.Constant{Inner: core.Allotment{*big.NewRat(125, 1000), *big.NewRat(375, 1000), *big.NewRat(1, 2)}},
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
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "EUR/2", Amount: core.NewAmountSpecific(99)}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob}},
			Error: "",
		},
	})
}

func TestSendAll(t *testing.T) {
	alice := core.Account("alice")
	bob := core.Account("bob")
	test(t, TestCase{
		Case: `send [EUR/2 *] (
	source = @alice
	destination = @bob
)`,
		Expected: CaseResult{
			Instructions: []byte{
				program.OP_APUSH, 00, 00,
				program.OP_APUSH, 01, 00,
				program.OP_IPUSH, 01, 00, 00, 00, 00, 00, 00, 00,
				program.OP_APUSH, 02, 00,
				program.OP_SEND,
			}, Resources: []program.Resource{
				program.Constant{Inner: core.Monetary{Asset: "EUR/2", Amount: core.NewAmountAll()}},
				program.Constant{Inner: alice},
				program.Constant{Inner: bob}},
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
