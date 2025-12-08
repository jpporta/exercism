package rotationalcipher

import (
	"strings"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"
const tolow = 'a' - 'A'

func RotationalCipher(plain string, shiftKey int) string {
	s := strings.Builder{}
	for _, c := range plain {
		// Not a letter
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			s.WriteRune(c)
			continue
		}
		diff := c-'A'		

		if c >= 'a' {
			diff -= tolow
		}
		t := alpha[(int(diff) + shiftKey) % 26] 

		if c >= 'a' {
			s.WriteByte(t)
		} else {
			s.WriteRune(rune(t) - tolow)
		}
	}
	return s.String()
}
