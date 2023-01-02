package core

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

const assetPattern = "^[A-Z/0-9]+$"

type Asset string

func (Asset) GetType() Type { return TypeAsset }
func (a Asset) String() string {
	return fmt.Sprintf("%v", string(a))
}

type HasAsset interface {
	GetAsset() Asset
}

func (a Asset) GetAsset() Asset { return a }

func ParseAsset(ass Asset) error {
	r, err := regexp.Compile(assetPattern)
	if err != nil {
		panic(errors.Wrap(err, "compiling assetPattern regexp"))
	}
	if !r.MatchString(string(ass)) {
		return fmt.Errorf("assets should respect pattern %s", assetPattern)
	}
	return nil
}
