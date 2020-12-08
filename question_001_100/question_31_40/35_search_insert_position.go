package question_31_40

// 35. 搜索插入位置
// https://leetcode-cn.com/problems/search-insert-position
// Topics: 数组 二分查找

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for right > left {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
