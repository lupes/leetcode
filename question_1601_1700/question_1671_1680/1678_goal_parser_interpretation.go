package question_1671_1680

// 1678. 设计 Goal 解析器
// https://leetcode-cn.com/problems/goal-parser-interpretation/
// Topics: 字符串

func interpret(command string) string {
	var res []byte
	for i := 0; i < len(command); {
		if command[i] == 'G' {
			res, i = append(res, 'G'), i+1
		} else if i+2 <= len(command) && command[i:i+2] == "()" {
			res, i = append(res, 'o'), i+2
		} else if i+4 <= len(command) && command[i:i+4] == "(al)" {
			res, i = append(res, 'a', 'l'), i+4
		}
	}
	return string(res)
}
