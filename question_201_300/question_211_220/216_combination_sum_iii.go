package question_211_220

// 216. 组合总和 III
// https://leetcode-cn.com/problems/combination-sum-iii/
// Topics: 数组 回溯算法

func combinationSum3(k int, n int) [][]int {
	if n > 45 || n < 1 || k > 9 || k < 1 {
		return nil
	}
	return combinationSum3Dfs(k, n, 1)
}

func combinationSum3Dfs(k int, n int, s int) [][]int {
	var res [][]int
	for i := s; i < 10 && i <= n; i++ {
		if k == 1 && i == n {
			res = append(res, []int{i})
		} else if k > 1 && n >= i*k+k-1 && n <= (k-1)*9+i {
			tmp := combinationSum3Dfs(k-1, n-i, i+1)
			for _, t := range tmp {
				res = append(res, append([]int{i}, t...))
			}
		}
	}
	return res
}
