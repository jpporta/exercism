package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	noCows   int
	errorMsg string
}

func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", ice.noCows, ice.errorMsg)
}

func DivideFood(fc FodderCalculator, noCows int) (float64, error) {
	total, err := fc.FodderAmount(noCows)
	if err != nil {
		return 0.0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return (total / float64(noCows)) * factor, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, noCows int) (float64, error) {
	if noCows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, noCows)
}

func ValidateNumberOfCows(noCows int) error {
	if noCows < 0 {
		return &InvalidCowsError{
			noCows:   noCows,
			errorMsg: "there are no negative cows",
		}
	}
	if noCows == 0 {
		return &InvalidCowsError{
			noCows:   noCows,
			errorMsg: "no cows don't need food",
		}
	}
	return nil

}
