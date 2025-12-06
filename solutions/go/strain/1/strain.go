package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func core[T any](list []T, fn func(T) bool, inverted bool) []T {
	nl := make([]T, 0)
	for _, el := range list {
		if fn(el) != inverted {
			nl = append(nl, el)
		}
	}
	return nl
}

func Keep[T any](list []T, fn func(T) bool) []T {
	return core(list, fn, false)
}

func Discard[T any](list []T, fn func(T) bool) []T {
	return core(list, fn, true)
}

