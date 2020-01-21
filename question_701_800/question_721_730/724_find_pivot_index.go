package question_721_730

// 724. 寻找数组的中心索引
// https://leetcode-cn.com/problems/find-pivot-index
// Topics: 数组

func pivotIndex(nums []int) int {
	var sum = 0
	for _, n := range nums {
		sum += n
	}
	var left, right = 0, sum
	for i, n := range nums {
		right -= n
		if left == right {
			return i
		}
		left += n
	}
	return -1
}
