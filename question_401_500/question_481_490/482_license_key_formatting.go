package question_481_490

// 482. 密钥格式化
// https://leetcode-cn.com/problems/license-key-formatting/
// Topics:

func licenseKeyFormatting(S string, K int) string {
	var t, r string
	for _, s := range S {
		if s >= 'a' && s <= 'z' {
			t += string(s - 'a' + 'A')
		} else if (s >= '0' && s <= '9') || (s >= 'A' && s <= 'Z') {
			t += string(s)
		}
	}
	l := len(t)
	if l == 0 {
		return ""
	}
	i := l % K
	if i != 0 {
		r = t[:i] + "-"
	}
	for k := 1; k <= (l-i)/K; k++ {
		r += t[(k-1)*K+i:k*K+i] + "-"
	}
	return r[:len(r)-1]
}
