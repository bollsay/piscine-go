package piscine

func Fibonacci(index int) func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
