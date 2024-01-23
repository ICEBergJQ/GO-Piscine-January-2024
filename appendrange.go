package piscine

func AppendRange(min, max int) []int {
	var rslt []int
	if min >= max {
		return nil
	}
	for min < max {
		rslt = append(rslt, min)
		min++
	}
	return rslt
}
