package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)
	for _, v := range runes {
		if !(v >= 65 && v <= 90) && !(v >= 97 && v <= 122) && !(v >= 48 && v <= 57) {
			return false
		}
	}
	return true
}
