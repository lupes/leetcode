package question_71_80

// 77. 组合
// https://leetcode-cn.com/problems/combinations
// Topics: 回溯算法

func combine(n int, k int) [][]int {
	return combineDfs(1, n, k)
}

func combineDfs(s, n, k int) [][]int {
	var res [][]int
	for i := s; i <= n; i++ {
		if k == 1 {
			res = append(res, []int{i})
		} else {
			for _, t := range combineDfs(i+1, n, k-1) {
				r := append([]int{i}, t...)
				res = append(res, r)
			}
		}
	}
	return res
}
