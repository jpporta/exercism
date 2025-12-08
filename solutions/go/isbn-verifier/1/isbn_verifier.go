package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	s := strings.ReplaceAll(isbn, "-", "")
	if len(s) != 10 {
		return false
	}
	sum := 0
	for i, ds := range s {
		d, err := strconv.Atoi(string(ds))
		if err != nil && ds == 'X' && i == len(s)-1 {
			err = nil
			d = 10
		}
		if err != nil {
			return false
		}
		sum += d * (10 - i)
	}
	return sum%11 == 0
}
