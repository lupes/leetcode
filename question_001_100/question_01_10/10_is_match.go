package question_01_10

// 10. 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	if (s == "" && p == "") || p == ".*" {
		return true
	}
	if s != "" && p == "" {
		return false
	}
	sLen, pLen := len(s), len(p)
	var t []string
	for i := 0; i < pLen; {
		if i < pLen-1 && p[i+1] == '*' {
			t = append(t, p[i:i+2])
			i += 2
		} else {
			t = append(t, p[i:i+1])
			i += 1
		}
	}
	tLen := len(t)
	flags := make([][]bool, tLen+1)
	for i := range flags {
		flags[i] = make([]bool, sLen+1)
	}
	flags[0][0] = true
	// 初始化零匹配
	for i, v := range t {
		if flags[i][0] && len(v) > 1 {
			flags[i+1][0] = true
		}
	}
	for i, v := range t {
		r := false
		for j, c := range s {
			// 零匹配
			if len(v) > 1 && flags[i][j+1] {
				r, flags[i+1][j+1] = true, true
			}
			// 多匹配
			if len(v) > 1 && (int32(v[0]) == c || v[0] == '.') && (flags[i][j] || flags[i+1][j]) {
				r, flags[i+1][j+1] = true, true
			}
			// 单匹配
			if (int32(v[0]) == c || v[0] == '.') && flags[i][j] {
				r, flags[i+1][j+1] = true, true
			}
		}
		if !r && !flags[i][0] {
			return false
		}
	}
	return flags[tLen][sLen]
}
