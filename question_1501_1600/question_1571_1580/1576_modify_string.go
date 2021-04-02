package question_1571_1580

// 1576. 替换所有的问号
// https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
// Topics: 字符串

func modifyString(s string) string {
	tmp := []byte("z" + s + "?")
	for i := 1; i < len(tmp)-1; i++ {
		if tmp[i] == '?' {
			t := (tmp[i-1] - 'a' + 1) % 26
			if t == tmp[i+1]-'a' {
				t = (t + 1) % 26
			}
			tmp[i] = t + 'a'
		}
	}
	return string(tmp[1 : len(tmp)-1])
}
