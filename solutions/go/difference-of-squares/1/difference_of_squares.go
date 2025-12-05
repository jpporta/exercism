package diffsquares

import "math"

func SquareOfSum(n int) int {
	return int(math.Pow(float64((n+1))*(float64(n)/2.0), 2))

}

func SumOfSquares(n int) int {
	sum := 0
	for i := range n {
		sum += int(math.Pow(float64(i+1), 2))
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
