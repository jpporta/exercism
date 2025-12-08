package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Cannot use negative or zero or > 64")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	var sum uint64 
	for i := range(64) {
		r, _ := Square(i+1)
		sum += r
	}
	return sum
}
