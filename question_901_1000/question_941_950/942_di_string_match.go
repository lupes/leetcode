package question_941_950

// 942. 增减字符串匹配
// https://leetcode-cn.com/problems/di-string-match
// Topics: 数学

func diStringMatch(S string) []int {
	res, left, right := make([]int, 0, len(S)+1), 0, len(S)
	for _, c := range S {
		if c == 'I' {
			res = append(res, left)
			left++
		} else {
			res = append(res, right)
			right--
		}
	}
	return append(res, left)
}
