package question_71_80

// 78. 子集
// https://leetcode-cn.com/problems/subsets

func subsets(nums []int) [][]int {
	var l = len(nums)
	var flag = make([]bool, l)
	var res [][]int
	for n := 0; n <= l/2; n++ {
		res = append(res, subsetsDfs(nums, flag, n, 0, l)...)
	}
	return res
}

func subsetsDfs(nums []int, flag []bool, n, i, l int) (res [][]int) {
	if n == 0 {
		return get(nums, flag)
	}
	last := l
	if l%2 == 0 && n == l/2 {
		last = 1
	}
	for j := i; j < last; j++ {
		flag[j] = true
		res = append(res, subsetsDfs(nums, flag, n-1, j+1, l)...)
		flag[j] = false
	}
	return
}

func get(nums []int, flag []bool) [][]int {
	var res = make([][]int, 2)
	for i, f := range flag {
		if f {
			res[0] = append(res[0], nums[i])
		} else {
			res[1] = append(res[1], nums[i])
		}
	}
	return res
}
