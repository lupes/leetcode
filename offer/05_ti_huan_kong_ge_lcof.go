package offer

// 剑指 Offer 05. 替换空格
// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// Topics: 字符串

func replaceSpace(s string) string {
	var res = make([]byte, 0, len(s))
	for i, c := range s {
		if c == ' ' {
			res = append(res, "%20"...)
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
