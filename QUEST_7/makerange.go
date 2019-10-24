package piscine

func MakeRange(min, max int) []int {
	if max > min {
		result := make([]int, max-min)
		i := 0
		for min < max {
			result[i] = min
			i++
			min++
		}
		return result
	}
	return nil
}
