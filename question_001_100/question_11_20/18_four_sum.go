package question_11_20

import "sort"

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for a := 0; a < size-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < size-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for c := b + 1; c < size-1; c++ {
				if c > b+1 && nums[c] == nums[c-1] {
					continue
				}
				for d := c + 1; d < size; d++ {
					if d > c+1 && nums[d] == nums[d-1] {
						continue
					}
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					}
				}
			}
		}
	}
	return res
}
