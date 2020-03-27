package question_901_910

// 908. 最小差值 I
// https://leetcode-cn.com/problems/smallest-range-i
// Topics: 数学

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, n := range A {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	if max-min < 2*K {
		return 0
	} else {
		return max - 2*K - min
	}
}
