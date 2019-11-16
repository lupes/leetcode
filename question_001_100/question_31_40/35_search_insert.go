package question_31_40

// 35. 搜索插入位置
// https://leetcode-cn.com/problems/search-insert-position

func searchInsert(nums []int, target int) int {
	i := -1
	num := 0
	for i, num = range nums {
		if num >= target {
			return i
		}
	}
	return i + 1
}
