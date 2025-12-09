package protein

import (
	"errors"
)

var ErrStop = errors.New("stop cordon found")
var ErrInvalidBase = errors.New("invalid cordon")

func FromRNA(rna string) ([]string, error) {
	res := make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
		if len(res) == 3 {
			return res, nil
		}
		switch rna[i : i+3] {
		case "AUG":
			res = append(res, "Methionine")
		case "UUU", "UUC":
			res = append(res, "Phenylalanine")
		case "UUA", "UUG":
			res = append(res, "Leucine")
		case "UCU", "UCC", "UCA", "UCG":
			res = append(res, "Serine")
		case "UAU", "UAC":
			res = append(res, "Tyrosine")
		case "UGU", "UGC":
			res = append(res, "Cysteine")
		case "UGG":
			res = append(res, "Tryptophan")
		case "UAA", "UAG", "UGA":
			if len(res) > 0 {
				return res, nil
			}
			return res, ErrStop
		default:
			return nil, ErrInvalidBase
		}
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	n, err := FromRNA(codon)
	if err != nil {
		return "", err
	}
	return n[0], nil
}
