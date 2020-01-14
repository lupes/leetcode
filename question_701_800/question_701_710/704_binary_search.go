package question_701_710

// 704. 二分查找
// https://leetcode-cn.com/problems/binary-search
// Topics: 二分查找

func search(nums []int, target int) int {
	left, right, center := 0, len(nums), 0
	for right > left {
		center = (left + right) / 2
		if nums[center] == target {
			return center
		} else if nums[center] > target {
			right = center
		} else {
			left = center + 1
		}
	}
	return -1
}
