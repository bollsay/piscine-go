package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for i, v := range runes {
		if i+1 == n {
			return v
		}
	}
	return '\x00'
}
