package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	used:= make(map[rune]struct{})
	for _, r := range strings.ToLower(word) {
		if int(r) < 97 || int(r) > 122 {
			continue
		}
		if _, has := used[r]; has {
			return false
		}
		used[r] = struct{}{}
	}
	return true
}
