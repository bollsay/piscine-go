package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i, v := range strs {
		if i != 0 {
			str += sep + v
		} else {
			str += v
		}
	}
	return str
}
