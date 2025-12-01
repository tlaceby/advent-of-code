package helpers

func Ternary[T any](cond bool, consequent, alternate T) T {
	if cond {
		return consequent
	}

	return alternate
}
