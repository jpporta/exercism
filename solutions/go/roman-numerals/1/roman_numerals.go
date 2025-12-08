package romannumerals

import (
	"errors"
	"strings"
)

type numeral struct {
	key   string
	value int
}

var numerals_table = []numeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ToRomanNumeral(input int) (string, error) {
	if input >= 4000 || input < 1 {
		return "", errors.New("only numbers between 4000 and 0")
	}
	n := input
	res := ""
	for _, num := range numerals_table {
		res += strings.Repeat(num.key, n/num.value)
		n -= num.value * (n / num.value)
	}
	return res, nil
}
