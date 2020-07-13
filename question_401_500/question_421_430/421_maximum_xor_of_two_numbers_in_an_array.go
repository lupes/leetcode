package question_421_430

// 421. 数组中两个数的最大异或值
// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
// Topics: 位运算 字典树

func findMaximumXOR(nums []int) int {
	var max int
	for i, a := range nums[:len(nums)-1] {
		for _, b := range nums[i+1:] {
			if a^b > max {
				max = a ^ b
			}
		}
	}
	return max
}
