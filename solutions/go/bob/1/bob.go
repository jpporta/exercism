package bob

import "strings"

func Hey(remark string) string {
	r := strings.TrimSpace(remark)
	if r == "" {
			return "Fine. Be that way!"	
	}

	is_l := r == strings.ToLower(r)
	is_u := r == strings.ToUpper(r)
	is_question := r[len(r) - 1] == '?'
	is_scream := is_u && !is_l
	
	switch {
		case is_question && is_scream:
			return "Calm down, I know what I'm doing!"
		case is_question:
			return "Sure."
		case is_scream:
			return "Whoa, chill out!"
		default:
		return "Whatever."
	}
}
