package question_41_50

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	tmp := permuteUniqueDfs(nums)
	m := make(map[string][]int)
	var res [][]int
	for _, r := range tmp {
		key := fmt.Sprintf("%v", r)
		_, ok := m[key]
		if !ok {
			m[key] = r
			res = append(res, r)
		}
	}
	return res
}

func permuteUniqueDfs(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	}
	for i, num := range nums {
		if i > 0 && (nums[i-1] == num || num == nums[0]) {
			continue
		}
		nums[0], nums[i] = nums[i], nums[0]
		for _, t := range permuteUniqueDfs(nums[1:]) {
			r := append([]int{num}, t...)
			res = append(res, r)
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res
}
