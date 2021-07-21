package core

import (
	"fmt"
	"math/big"
	"regexp"
)

func ParsePortion(input string) (*Portion, bool) {
	var res *big.Rat

	re := regexp.MustCompile(`^([0-9]+)(?:[.]([0-9]+))?[%]$`)
	percent_match := re.FindStringSubmatch(input)
	fmt.Println(input)
	fmt.Println(percent_match)
	if len(percent_match) != 0 {
		integral := percent_match[1]
		fractional := percent_match[2]
		rat, ok := new(big.Rat).SetString(integral + "." + fractional)
		if !ok {
			return nil, false
		}
		rat.Mul(rat, big.NewRat(1, 100))
		res = rat
	} else {
		re = regexp.MustCompile(`^([0-9]+)\s?[/]\s?([0-9]+)$`)
		fraction_match := re.FindStringSubmatch(input)
		if len(fraction_match) != 0 {
			numerator := fraction_match[1]
			denominator := fraction_match[2]
			rat, ok := new(big.Rat).SetString(numerator + "/" + denominator)
			if !ok {
				return nil, false
			}
			res = rat
		}
	}

	if res == nil || res.Cmp(big.NewRat(0, 1)) != 1 || res.Cmp(big.NewRat(1, 1)) != -1 {
		return nil, false
	}
	portion := Portion(*res)
	return &portion, true
}
