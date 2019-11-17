package question_261_270

// 268. 缺失数字
// https://leetcode-cn.com/problems/missing-number/
// Topics: 位运算 数组 数学

func missingNumber(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res ^= n ^ i
	}
	return res
}
