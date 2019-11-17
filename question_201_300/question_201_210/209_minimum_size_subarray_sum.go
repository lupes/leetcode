package question_201_210

// 209. 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// Topics: 数组 双指针 二分查找

func minSubArrayLen(s int, nums []int) int {
	min := len(nums) + 1
	for i, _ := range nums {
		t := 0
		for j := i; j < len(nums); j++ {
			t += nums[j]
			if t >= s {
				if j-i+1 < min {
					min = j - i + 1
				}
				break
			}
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
