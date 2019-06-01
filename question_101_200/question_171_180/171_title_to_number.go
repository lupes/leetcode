package question_171_180

import "math"

func titleToNumber(s string) int {
	var res int
	size := len(s)
	for i, s := range s {
		res += int(s-'A'+1) * int(math.Pow(26, float64(size-i-1)))
	}
	return res
}
