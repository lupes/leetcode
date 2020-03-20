package question_921_930

// 929. 独特的电子邮件地址
// https://leetcode-cn.com/problems/unique-email-addresses
// Topics: 字符串

func numUniqueEmails(emails []string) int {
	var flag = make(map[string]struct{}, 0)
	for _, email := range emails {
		var res, t = "", true
		for i, c := range email {
			if c == '@' {
				res += email[i:]
				flag[res] = struct{}{}
			}
			if c == '+' {
				t = false
			}
			if t {
				if c != '.' {
					res += email[i : i+1]
				}
			}
		}
	}
	return len(flag)
}
