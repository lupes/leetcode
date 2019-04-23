package question_41_50

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	flag := make([][]bool, pLen+1)
	for i := range flag {
		flag[i] = make([]bool, sLen+1)
	}
	flag[0][0] = true
	for i := range p {
		if p[i] == '*' && flag[i][0] {
			flag[i+1][0] = true
		}
		for j := range s {
			if flag[i][j] && (p[i] == s[j] || p[i] == '?') {
				flag[i+1][j+1] = true
			} else if p[i] == '*' && (flag[i][j] || flag[i][j+1] || flag[i+1][j]) {
				flag[i+1][j+1] = true
			}
		}
	}
	return flag[pLen][sLen]
}
