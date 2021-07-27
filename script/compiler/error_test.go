package compiler

import (
	"fmt"
	"testing"
)

func TestEndCharacter(t *testing.T) {
	src := `
	send [CREDIT 200] (
		source = @a
		destination = {
			500% to @b
			50% to @c
		}
	)
	`

	_, err := Compile(src)
	if err == nil {
		t.Fatal("expected error and got none")
	}

	if _, ok := err.(*CompileErrorList); !ok {
		t.Fatal("error had wrong type")
	}

	lerr := err.(*CompileErrorList).errors[0]

	if lerr.startl != 5 {
		t.Fatal(fmt.Sprintf("start line was %v", lerr.startl))
	}
	if lerr.startc != 3 {
		t.Fatal(fmt.Sprintf("start character was %v", lerr.startc))
	}
	if lerr.endl != 5 {
		t.Fatal(fmt.Sprintf("end line was %v", lerr.endl))
	}
	if lerr.endc != 7 {
		t.Fatal(fmt.Sprintf("end character was %v", lerr.endc))
	}
}
