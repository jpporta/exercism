package proverb

import (
	"fmt"
)

func Proverb(rhyme []string) []string {
	rhymes := make([]string, 0, len(rhyme))
	for i, cur := range rhyme {
		if i == len(rhyme) - 1 {
			rhymes = append(rhymes, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			continue
		}
		rhymes = append(rhymes, fmt.Sprintf("For want of a %s the %s was lost.", cur, rhyme[i+1]))
	}
	return rhymes
}
