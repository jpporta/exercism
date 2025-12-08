package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
)

func KindFromSides(a, b, c float64) Kind {
	switch {
	case a == b && b == c && a > 0:
		return Equ
	case a >= b + c || b >= a + c || c >= a + b:
		return NaT
	case a != b && b != c && c != a:
		return Sca
	default:
		return Iso
	}
}
