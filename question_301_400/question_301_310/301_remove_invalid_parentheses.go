package question_301_310

// 301. 删除无效的括号
// https://leetcode-cn.com/problems/remove-invalid-parentheses
// Topics: 深度优先搜索 广度优先搜索

func removeInvalidParentheses(s string) []string {
	if invalidParentheses(s) {
		return []string{s}
	}

	var now = map[string]struct{}{s: {}}
	var valid = make(map[string]struct{}, 0)
	for len(now) > 0 {
		var next = make(map[string]struct{}, 0)
		for tmp := range now {
			for i, c := range tmp {
				if c == ')' || c == '(' {
					t := tmp[:i] + tmp[i+1:]
					if invalidParentheses(t) {
						valid[t] = struct{}{}
					} else {
						next[t] = struct{}{}
					}
				}
			}
		}
		if len(valid) > 0 {
			break
		}
		now = next
	}

	var res []string
	for k := range valid {
		res = append(res, k)
	}

	return res
}

func invalidParentheses(s string) bool {
	var a = 0
	for _, c := range s {
		if c == '(' {
			a++
		} else if c == ')' {
			a--
		}
		if a < 0 {
			return false
		}
	}
	return a == 0
}
