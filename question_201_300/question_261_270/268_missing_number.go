package question_261_270

// 268. 缺失数字
// https://leetcode-cn.com/problems/missing-number/

func missingNumber(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res ^= n ^ i
	}
	return res
}
