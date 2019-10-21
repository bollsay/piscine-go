package piscine

func Fibonacci(index int) func() int {
	if index < 0 {
		return -1
	}
	if index == 0 || index == 1 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
