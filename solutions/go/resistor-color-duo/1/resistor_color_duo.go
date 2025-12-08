package resistorcolorduo

import "math"

var colorCodes = map[string]int{
"black": 0,
"brown": 1,
"red": 2,
"orange": 3,
"yellow": 4,
"green": 5,
"blue": 6,
"violet": 7,
"grey": 8,
"white": 9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	num := 0
	cs := colors[:2]
	for i, c := range cs {
		if _,  ok := colorCodes[c]; !ok {
			continue
		}
		num += colorCodes[c] * int(math.Pow10((len(cs) - i - 1))) 
	}

	return num
}
