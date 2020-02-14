package question_821_830

// 824. 山羊拉丁文
// https://leetcode-cn.com/problems/goat-latin
// Topics: 字符串

func toGoatLatin(S string) string {
	S += " "
	var res, end, start = "", "maa", -1
	for i, c := range S {
		if c != ' ' && start == -1 {
			start = i
		} else if c == ' ' && start != -1 {
			if S[start] == 'a' || S[start] == 'e' || S[start] == 'i' || S[start] == 'o' || S[start] == 'u' ||
				S[start] == 'A' || S[start] == 'E' || S[start] == 'I' || S[start] == 'O' || S[start] == 'U' {
				res += " " + S[start:i] + end
			} else {
				res += " " + S[start+1:i] + S[start:start+1] + end
			}
			end += "a"
			start = -1
		}
	}
	return res[1:]
}
