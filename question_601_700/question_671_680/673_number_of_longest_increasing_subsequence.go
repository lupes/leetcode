package question_671_680

// 673. 最长递增子序列的个数
// https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence
// Topics: 动态规划

func findNumberOfLIS(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	var res, max, lengths, counts = 0, 0, make([]int, l), make([]int, l)
	for i, n := range nums {
		counts[i] = 1
		for j, t := range nums[:i] {
			if n > t {
				if lengths[j] >= lengths[i] {
					lengths[i] = 1 + lengths[j]
					counts[i] = counts[j]
					if lengths[i] > max {
						max = lengths[i]
					}
				} else if lengths[j]+1 == lengths[i] {
					counts[i] += counts[j]
				}
			}
		}
	}
	for i, n := range counts {
		if lengths[i] == max {
			res += n
		}
	}
	return res
}
