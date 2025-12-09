package anagram

import (
	"sort"
	"strings"
)

type Runes []rune

func (r Runes) Len() int           { return len(r) }
func (r Runes) Less(a, b int) bool { return r[a] < r[b] }
func (r Runes) Swap(a, b int)      { r[a], r[b] = r[b], r[a] }
func (r Runes) Eq(o Runes) bool {
	if r.Len() != o.Len() {
		return false
	}
	for i := range r {
		if r[i] != o[i] {
			return false
		}
	}
	return true
}

func Detect(subject string, candidates []string) []string {
	sub := strings.ToLower(subject)
	n := Runes(sub)
	sort.Sort(n)
	res := make([]string, 0)
	for _, cr := range candidates {
		cs := strings.ToLower(cr)
		if sub == cs {
			continue
		}
		c := Runes(cs)
		sort.Sort(c)
		if n.Eq(c) {
			res = append(res, cr)
		}
	}
	return res
}
