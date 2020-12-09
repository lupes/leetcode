package question_0011_0020

// 26. 删除排序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
// Topics: 数组 双指针

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
