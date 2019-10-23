package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for _, v := range runes {
		if !(v >= 32 && v <= 127) {
			return false
		}
	}
	return true
}
