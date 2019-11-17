package question_11_20

import "sort"

// 15. 三数之和
// https://leetcode-cn.com/problems/3sum
// Topics: 数组 双指针

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	if size < 3 {
		return [][]int{}
	}
	if nums[0] == 0 && nums[size-1] == 0 {
		return [][]int{{0, 0, 0}}
	}
	var res [][]int
	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		l, r := i+1, size-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < size-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			tmp := nums[l] + nums[r]
			if tmp == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				r--
				l++
			} else if tmp > target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
