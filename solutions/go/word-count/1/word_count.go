package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	r := regexp.MustCompile(`\w+'?-?\w+|\d|[a-zA-Z]`)
	ws := r.FindAllString(phrase, -1)
	f := make(Frequency, 0)
	for _, w := range ws {
		if w == "" {
			continue
		}
		f[strings.ToLower(w)]++
	}
	return f
}
