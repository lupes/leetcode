package question_991_1000

// 997. 找到小镇的法官
// https://leetcode-cn.com/problems/find-the-town-judge
// Topics: 图

func findJudge(N int, trust [][]int) int {
	var flag = make([][2]int, N+1)
	for _, tmp := range trust {
		flag[tmp[1]][0]++
		flag[tmp[0]][1]++
	}
	for k, v := range flag[1:] {
		if v[0] == N-1 && v[1] == 0 {
			return k + 1
		}
	}
	return -1
}
