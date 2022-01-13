package question_741_750

// 747. 至少是其他数字两倍的最大数
// https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
// Topics: 数组

func dominantIndex(nums []int) int {
	a, b, i := -1, -1, 0
	for j, n := range nums {
		if n > a {
			a, b, i = n, a, j
		} else if n > b {
			b = n
		}
	}
	if b*2 <= a {
		return i
	}
	return -1
}
