package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	q := []rune(toFind)

	for i, v := range runes {
		if v == q[0] {
			found := true
			for j, qv := range q {
				if runes[i+j] != qv {
					found = false
				}
			}
			if found == true {
				return i
			}
		}
	}
	return -1
}
