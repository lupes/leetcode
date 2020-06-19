package question_1261_1270

// 1267. 统计参与通信的服务器
// https://leetcode-cn.com/problems/count-servers-that-communicate/
// Topics: 图 数组

func countServers(grid [][]int) int {
	var res, row, col = 0, make(map[int]struct{}), make(map[int]struct{})
	for i, r := range grid {
		for j, c := range r {
			if c == 1 {
				f := false
				if _, ok := row[i]; !ok {
					for k := j + 1; k < len(r); k++ {
						if r[k] == 1 {
							res++
							f = true
							break
						}
					}
				} else {
					f = true
					res++
				}
				if _, ok := col[j]; !ok && !f {
					for k := i + 1; k < len(grid); k++ {
						if grid[k][j] == 1 {
							res++
							break
						}
					}
				} else if !f {
					res++
				}
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}
	return res
}
