package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	is := []rune(toFind)
	counter := 0
	ip := 0
	for range runes {
		counter++
	}
	for range is {
		ip++
	}

	for i, v := range runes {
		if v == is[0] {
			found := true
			for j, qv := range is {
				if counter-1-i >= ip {
					if runes[i+j] != qv {
						found = false
					}
				} else {
					return -1
				}
			}
			if found == true {
				return i
			}
		}
	}
	return -1
}
