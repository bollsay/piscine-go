package piscine

func AppendRange(min, max int) []int {
	if max > min {
		var result []int
		for min < max {
			result = append(result, min)
			min++
		}
		return result
	}
	return nil
}
