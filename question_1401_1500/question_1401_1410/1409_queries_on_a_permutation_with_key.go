package question_1401_1410

// 1409. 查询带键的排列
// https://leetcode-cn.com/problems/queries-on-a-permutation-with-key/
// Topics: 数组

func processQueries(queries []int, m int) []int {
	var flag = make([]int, m+1)
	for i := 1; i <= m; i++ {
		flag[i] = i
	}

	var res = make([]int, 0, len(queries))
	for _, n := range queries {
		var t = flag[0]
		for j := 1; j <= m; j++ {
			flag[j], t = t, flag[j]
			if t == n {
				res = append(res, j-1)
				break
			}
		}
		flag[1] = n
	}
	return res
}
