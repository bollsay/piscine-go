package piscine

func Split(str, charset string) []string {
	counter := 0
	for range charset {
		counter++
	}
	runes := []rune(str)
	rCounter := 0
	for range runes {
		rCounter++
	}

	chRunes := []rune(charset)
	ch := 0
	checker := make([]rune, counter)
	for i := 0; i < rCounter-counter; i++ {
		for j := 0; j < counter; j++ {
			checker[j] = runes[i+j]
		}
		res := true
		for k := 0; k < counter; k++ {
			if checker[k] != chRunes[k] {
				res = false
			}
		}
		if res {
			ch++
		}
	}
	result := make([]string, ch+1)
	dot := 0
	x := 0
	for i := 0; i < rCounter-counter; i++ {
		for j := 0; j < counter; j++ {
			checker[j] = runes[i+j]
		}
		res := true
		for k := 0; k < counter; k++ {
			if checker[k] != chRunes[k] {
				res = false
			}
		}
		if res {
			result[x] = string(str[dot:i])
			x++
			dot = i + counter
		}
	}
	result[x] = string(str[dot:])
	return result
}
