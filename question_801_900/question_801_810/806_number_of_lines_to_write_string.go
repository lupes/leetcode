package question_801_810

// 806. 写字符串需要的行数
// https://leetcode-cn.com/problems/number-of-lines-to-write-string
// Topics:

func numberOfLines(widths []int, S string) []int {
	var res = []int{1, 0}
	for _, c := range S {
		if res[1]+widths[c-'a'] > 100 {
			res[0]++
			res[1] = 0
		}
		res[1] += widths[c-'a']
	}
	return res
}
