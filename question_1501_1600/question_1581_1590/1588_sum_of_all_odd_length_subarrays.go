package question_1581_1590

// 1588. 所有奇数长度子数组的和
// https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
// Topics: 数组 前缀和

func sumOddLengthSubarrays(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	nums = append([]int{0}, nums...)
	var res int
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j -= 2 {
			res += nums[i] - nums[j]
		}
	}
	return res
}
