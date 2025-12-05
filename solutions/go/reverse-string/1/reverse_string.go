package reverse

func Reverse(input string) string {
	var rev string
	for _, r := range input {
		rev = string(r) + rev
	}
	return rev
}
