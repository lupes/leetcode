package question_821_830

// 821. 字符的最短距离
// https://leetcode-cn.com/problems/shortest-distance-to-a-character
// Topics:

func shortestToChar(S string, C byte) []int {
	var tmp []int
	for i, c := range S {
		if c == int32(C) {
			tmp = append(tmp, i)
		}
	}
	var res []int
	var last = -1
	for i, c := range S {
		var t = 10001
		if c != int32(C) {
			if last < len(tmp)-1 && tmp[last+1]-i < t {
				t = tmp[last+1] - i
			}
			if last > -1 && i-tmp[last] < t {
				t = i - tmp[last]
			}
		} else {
			t = 0
			last++
		}
		res = append(res, t)
	}
	return res
}
