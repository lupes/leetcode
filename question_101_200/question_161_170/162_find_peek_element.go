package question_161_170

// 162. 寻找峰值
// https://leetcode-cn.com/problems/find-peak-element

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for r > l {
		c := (l + r) / 2
		if c == l {
			return c
		}
		if nums[c] < nums[c-1] {
			r = c
		} else {
			l = c
		}
	}
	return -1
}
