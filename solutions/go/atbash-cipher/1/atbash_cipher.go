package atbash

import "strings"

var alpha = "abcdefghijklmnopqrstuvwxyz"

func Atbash(s string) string {
	sb := strings.Builder{}
	for _, r := range strings.ToLower(s) {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') {
			continue
		}
		if r >= '0' && r <= '9' {
			sb.WriteRune(r)
		} else {
			sb.WriteByte(alpha[26-(r-'a')-1])
		}
		if sb.Len()%6 == 5 {
			sb.WriteRune(' ')
		}
	}
	return strings.TrimSpace(sb.String())
}
