package piscine

func IsUpper(str string) bool {
	runes := []rune(str)
	for _, v := range runes {
		if !(v >= 65 && v <= 90) {
			return false
		}
	}
	return true
}
