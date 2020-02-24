package question_821_830

// 821. 字符的最短距离
// https://leetcode-cn.com/problems/shortest-distance-to-a-character
// Topics:

func shortestToChar(S string, C byte) []int {
	var last = -1
	var res = make([]int, len(S))
	for i := range S {
		var t = 10001
		if S[i] != C && last != -1 {
			t = i - last
		} else if S[i] == C {
			t = 0
			last = i
		}
		res[i] = t
	}
	last = -1
	for i := len(S) - 1; i >= 0; i-- {
		var t = 10001
		if S[i] != C && last != -1 {
			t = last - i
		} else if S[i] == C {
			t = 0
			last = i
		}
		if t < res[i] {
			res[i] = t
		}
	}
	return res
}
