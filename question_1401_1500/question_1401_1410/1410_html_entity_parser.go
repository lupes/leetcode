package question_1401_1410

// 1410. HTML 实体解析器
// https://leetcode-cn.com/problems/html-entity-parser/
// Topics: 栈 字符串

var flag = map[string]byte{"&quot;": '"', "&apos;": '\'', "&amp;": '&', "&gt;": '>', "&lt;": '<', "&frasl;": '/'}

func entityParser(text string) string {
	l, start, kl, res := len(text), len(text), 0, make([]byte, len(text))
	for i := l - 1; i >= 0; i-- {
		start--
		res[start] = text[i]
		if text[i] == '&' {
			for k, v := range flag {
				kl = len(k)
				if kl <= l-start && k == string(res[start:start+kl]) {
					start += kl - 1
					res[start] = v
					break
				}
			}
		}
	}
	return string(res[start:])
}
