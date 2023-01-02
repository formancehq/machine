package core

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

const accountPattern = "^[a-zA-Z_]+[a-zA-Z0-9_:]*$"

type Account string

func (Account) GetType() Type { return TypeAccount }
func (a Account) String() string {
	return fmt.Sprintf("@%v", string(a))
}

func ParseAccount(acc Account) error {
	r, err := regexp.Compile(accountPattern)
	if err != nil {
		panic(errors.Wrap(err, "compiling accountPattern regexp"))
	}
	if !r.MatchString(string(acc)) {
		return fmt.Errorf("accounts should respect pattern %s", accountPattern)
	}
	return nil
}
