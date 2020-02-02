package question_741_750

// 747. 至少是其他数字两倍的最大数
// https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
// Topics: 数组

func dominantIndex(nums []int) int {
	max1, index1, max2 := -1, -1, -1
	for i, n := range nums {
		if n > max1 {
			max1, max2 = n, max1
			index1 = i
		} else if n > max2 {
			max2 = n
		}
	}
	if (max2 == 0 && max1 > 0) || (max2 > 0 && max1/max2 >= 2) || (len(nums) == 1) {
		return index1
	}
	return -1
}
