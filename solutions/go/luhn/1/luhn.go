package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	should_double := false
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
		n, err := strconv.Atoi(string(id[i]))
		if err != nil {
			continue
		}
		if should_double {
			should_double = false
			double := 2 * n
			if double > 9 {
				double -= 9
			}
			sum += double
			continue
		}
		should_double = true
		sum += n
	}
	return sum % 10 == 0
}
