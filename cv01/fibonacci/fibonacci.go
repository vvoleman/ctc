package fibonacci

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		val := a

		// No need for swap variable
		a, b = b, a+b

		return val
	}
}
