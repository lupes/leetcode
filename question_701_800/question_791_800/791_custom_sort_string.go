package question_791_800

// 791. 自定义字符串排序
// https://leetcode-cn.com/problems/custom-sort-string
// Topics: 字符串

func customSortString(S string, T string) string {
	var tmp = make(map[int32]string)
	for _, c := range T {
		tmp[c] += string(c)
	}
	var res = ""
	for _, c := range S {
		res += tmp[c]
		delete(tmp, c)
	}
	for _, v := range tmp {
		res += v
	}
	return res
}
