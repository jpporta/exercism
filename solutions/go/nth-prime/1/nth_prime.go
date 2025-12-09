package prime

import "errors"

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid n, must be >= 1")
	}

	found := 0
	for i := 2; found < n; i++ {
		if isPrime(i) {
			found++
			if found == n {
				return i, nil
			}
		}
	}
	return 0, errors.New("you shouldn't have reached this lol")
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
