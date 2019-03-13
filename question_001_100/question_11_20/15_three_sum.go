package question_11_20

import "sort"

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
		for j := i + 1; j < size; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < size; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}
