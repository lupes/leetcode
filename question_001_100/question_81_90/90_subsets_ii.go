package question_81_90

import "sort"

// 90. 子集 II
// https://leetcode-cn.com/problems/subsets-ii

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	flag := make([]bool, n)
	var res [][]int
	for i := 0; i <= n; i++ {
		t := subsetsWithDupDfs(nums, flag, 0, i)
		res = append(res, t...)
	}
	return res
}

func subsetsWithDupDfs(nums []int, flag []bool, s, n int) (res [][]int) {
	if n == 0 {
		res = append(res, getSubsets(nums, flag))
		return res
	}
	for i := s; i <= len(nums)-n; i++ {
		if i != s && nums[i] == nums[i-1] {
			continue
		}
		flag[i] = true
		t := subsetsWithDupDfs(nums, flag, i+1, n-1)
		res = append(res, t...)
		flag[i] = false
	}
	return res
}

func getSubsets(nums []int, flag []bool) []int {
	res := []int{}
	for i, t := range flag {
		if t {
			res = append(res, nums[i])
		}
	}
	return res
}
