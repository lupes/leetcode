package question_71_80

// 80. 删除排序数组中的重复项 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii

func removeDuplicates(nums []int) int {
	now, l, flag := 0, len(nums), false
	if l <= 2 {
		return l
	}
	for i := 1; i < l; i++ {
		if nums[i] == nums[now] {
			if flag {
				continue
			} else {
				flag = true
			}
		} else {
			flag = false
		}
		now++
		nums[now] = nums[i]
	}
	return now + 1
}
