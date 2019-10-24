package piscine

func AlphaCount(str string) int {
	strings := []rune(str)
	counter := 0
	for _, v := range strings {
		if (v >= rune(65) && v <= rune(90)) || (v >= rune(97) && v <= rune(122)) {
			counter++
		}
	}
	return counter
}
