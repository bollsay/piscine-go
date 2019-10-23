package piscine

func StrLen(str string) int {
	nc := 0
	for range str {
		nc = nc + 1
	}
	return nc
}
