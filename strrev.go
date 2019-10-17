package piscine

func StrRev(s string) string {
	r := []rune(s)
	nc := 0
	nword := ""
	for range s {
		nc = nc + 1
	}
	for i := nc - 1; i >= 0; i-- {
		nword = nword + string(r[i])
	}
	return string(nword)
}
