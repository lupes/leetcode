package question_821_830

// 830. 较大分组的位置
// https://leetcode-cn.com/problems/positions-of-large-groups
// Topics: 数组

func largeGroupPositions(S string) [][]int {
	var res [][]int
	start := 0
	for i := 1; i < len(S); i++ {
		if S[i] != S[start] {
			if i-start >= 3 {
				res = append(res, []int{start, i - 1})
			}
			start = i
		}
	}
	if len(S)-start >= 3 {
		res = append(res, []int{start, len(S) - 1})
	}
	return res
}
