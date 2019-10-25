package question_531_540

// 540. 有序数组中的单一元素
// https://leetcode-cn.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
	l, r, h, c, n := 0, len(nums), len(nums)-1, 0, 0
	if h == 0 || nums[0] != nums[1] {
		return nums[0]
	}
	if nums[h] != nums[h-1] {
		return nums[h]
	}
	for l <= r {
		c = (l + r) / 2
		if (c > 0 && c < h) && (nums[c] != nums[c-1] && nums[c] != nums[c+1]) {
			return nums[c]
		} else if c > 0 && c < h {
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
		} else if c == h {
			r = c - 2
		}
	}
	return -1
}
