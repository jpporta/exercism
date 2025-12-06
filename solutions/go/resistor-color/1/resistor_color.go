package resistorcolor

var COLORS = [10]string{ 
"black",
"brown",
"red",
"orange",
"yellow",
"green",
"blue",
"violet",
"grey",
"white",
 }

func Colors() []string {
	return COLORS[:]
}

func ColorCode(color string) int {
	for v, c := range COLORS {
		if c == color {
			return v
		}
	}
	return 0
}
