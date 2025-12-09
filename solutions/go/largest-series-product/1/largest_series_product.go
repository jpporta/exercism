package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("span bigger than input")
	}
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	var big int64
	for i := 0; i <= len(digits)-span; i++ {
		var cur int64 = 1
		for j := range span {
			if digits[i+j] < '0' || digits[i+j] > '9' {
				return 0, errors.New("digits input must only contain digits")
			}
			cur *= int64(digits[i+j] - '0')
		}
		if cur > big {
			big = cur
		}
	}
	return big, nil
}
