package question_531_540

// 540. 有序数组中的单一元素
// https://leetcode-cn.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
	l, r, h, c := 0, len(nums), len(nums), 0
	for l <= r {
		c = (l + r) / 2
		if (c == 0 && c == h-1) || (c == 0 && nums[c] != nums[c+1]) || (c == h-1 && nums[c] != nums[c-1]) ||
			((c > 0 && c < h-1) && (nums[c] != nums[c-1] && nums[c] != nums[c+1])) {
			return nums[c]
		}
		if c > 0 && c < h-1 {
			n := 0
			if nums[c] == nums[c-1] {
				n = c
			} else if nums[c] == nums[c+1] {
				n = c + 1
			}
			if n%2 == 0 {
				r = n - 2
			} else {
				l = n + 1
			}
		} else if c == 0 {
			l = c + 2
		} else if c == h-1 {
			r = c - 2
		}
	}
	return -1
}
