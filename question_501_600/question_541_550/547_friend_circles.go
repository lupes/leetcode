package question_541_550

// 547. 朋友圈
// https://leetcode-cn.com/problems/friend-circles
// Topics: 深度优先搜索 并查集

func findCircleNum(M [][]int) int {
	var res = make(map[int]map[int]struct{}, 0)
	for i := 0; i < len(M); i++ {
		res[i] = make(map[int]struct{})
		res[i][i] = struct{}{}
	}
	for i, row := range M {
		for j := i + 1; j < len(row); j++ {
			if M[i][j] == 0 {
				continue
			}
			var ik, jk = -1, -1
			for k, flag := range res {
				if _, ok := flag[i]; ok {
					flag[j] = struct{}{}
					ik = k
				} else if _, ok := flag[j]; ok {
					flag[i] = struct{}{}
					jk = k
				}
			}

			if ik != -1 && jk != -1 && ik != jk {
				for n, _ := range res[jk] {
					res[ik][n] = struct{}{}
				}
				delete(res, jk)
			}
		}
	}
	return len(res)
}
