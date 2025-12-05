package dna

import (
	"errors"
)

type Histogram map[rune]int
type DNA []rune

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _,r := range d {
		switch r {
		case 'A', 'C', 'G', 'T':
			h[r]++
		default:
			return h, errors.New("invalid dna")
		}
	}
	return h, nil
}
