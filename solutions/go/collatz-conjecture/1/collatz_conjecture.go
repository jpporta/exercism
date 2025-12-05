package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return n, errors.New("Can't be 0 or negative")
	}
	if n == 1 {
		return 0, nil
	}
	nextStep := 0
	if n%2 == 0 {
		nextStep = n / 2
	} else {
		nextStep = 3*n + 1
	}
	res, err := CollatzConjecture(nextStep)
	if err != nil {
		return res, err
	}
	return res + 1, nil
}
