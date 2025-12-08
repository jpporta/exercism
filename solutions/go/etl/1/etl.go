package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, 26)
	for i := range in {
		for _, l := range in[i] {
			out[strings.ToLower(l)] = i
		}
	}
	return out
}
