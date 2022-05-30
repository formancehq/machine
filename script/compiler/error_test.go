package compiler

import (
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

	lerr := err.(*CompileErrorList).Errors[0]

	if lerr.Startl != 5 {
		t.Fatalf("start line was %v", lerr.Startl)
	}
	if lerr.Startc != 3 {
		t.Fatalf("start character was %v", lerr.Startc)
	}
	if lerr.Endl != 5 {
		t.Fatalf("end line was %v", lerr.Endl)
	}
	if lerr.Endc != 7 {
		t.Fatalf("end character was %v", lerr.Endc)
	}
}
