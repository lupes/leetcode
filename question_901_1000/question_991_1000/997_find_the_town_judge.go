package question_991_1000

// 997. 找到小镇的法官
// https://leetcode-cn.com/problems/find-the-town-judge
// Topics: 图

func findJudge(N int, trust [][]int) int {
	if N == 1 {
		return 1
	}
	var flagA, flagB = make(map[int][]int), make(map[int]struct{})
	for _, tmp := range trust {
		flagA[tmp[1]] = append(flagA[tmp[1]], tmp[0])
		flagB[tmp[0]] = struct{}{}
	}
	for k, v := range flagA {
		if _, ok := flagB[k]; len(v) == N-1 && !ok {
			return k
		}
	}
	return -1
}
