package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	for _, v := range runes {
		if !(v >= 97 && v <= 122) {
			return false
		}
	}
	return true
}
