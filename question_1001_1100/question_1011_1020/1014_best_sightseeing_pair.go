package question_1011_1020

// 1014. 最佳观光组合
// https://leetcode-cn.com/problems/best-sightseeing-pair
// Topics: 数组

func maxScoreSightseeingPair(A []int) int {
	var max = 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < i+1000 && j < len(A); j++ {
			t := A[i] + A[j] + i - j
			if t > max {
				max = t
			}
		}
	}
	return max
}
