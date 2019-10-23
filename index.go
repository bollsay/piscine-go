package piscine

func Index(s string, toFind string) int {
	runes := []rune(s)
	q := []rune(toFind)
	counter := 0
	qc := 0
	for range runes {
		counter++
	}
	for range q {
		qc++
	}

	for i, v := range runes {
		if v == q[0] {
			found := true
			for j, qv := range q {
				if runes[i+j] != qv {
					found = false
					if counter-1-i >= qc {
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
}
