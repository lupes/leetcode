package question_231_240

// 238. 除自身以外数组的乘积
// https://leetcode-cn.com/problems/product-of-array-except-self/
// Topics: 数组

func productExceptSelf(nums []int) []int {
	var res = make([]int, len(nums))
	left, right := 1, 1
	for i, n := range nums {
		left = n * left
		res[i] = left
	}
	for i := len(nums) - 2; i >= 0; i-- {
		res[i+1] = res[i] * right
		right *= nums[i+1]
	}
	res[0] = right
	return res
}
