package question_81_90

// 81. 搜索旋转排序数组 II
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
// Topics: 数组 二分查找

func search(nums []int, target int) bool {
	for _, n := range nums {
		if target == n {
			return true
		}
	}
	return false
}
