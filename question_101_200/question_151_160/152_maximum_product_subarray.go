package question_151_160

// 152. 乘积最大子序列
// https://leetcode-cn.com/problems/maximum-product-subarray/
// Topics: 数组 动态规划

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i, _ := range nums {
		t := 1
		for j := i; j < len(nums); j++ {
			t *= nums[j]
			if t > max {
				max = t
			}
		}
	}
	return max
}
