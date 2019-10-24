package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, v := range args {
		if i == 0 {
			str += v
		} else {
			str += "\n" + v
		}
	}
	return str
}
