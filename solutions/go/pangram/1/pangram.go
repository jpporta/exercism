package pangram

func IsPangram(input string) bool {
	var checks [26]bool
	for _, r := range input {
		var idx rune
		if r >= 'a' {
			idx = r - 'a'
		} else {
			idx = r - 'A'
		}
		if idx < 0 || idx > 26 {
			continue
		}
		checks[idx] = true
	}
	for _, c := range checks {
		if !c {
			return false
		}
	}
	return true
}
