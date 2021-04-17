package question_1611_1620

// 1616. 分割两个字符串得到回文串
// https://leetcode-cn.com/problems/split-two-strings-to-make-palindrome/
// Topics: 贪心算法 双指针 字符串

func checkPalindromeFormation(a string, b string) bool {
	return checkPalindromeFormationHelper(a, b) || checkPalindromeFormationHelper(b, a)
}

func checkPalindromeFormationHelper(a, b string) bool {
	l, lt, r, rt, t := 0, 0, len(a)-1, len(a)-1, 0
	for r > l {
		//if t == 0 {
		//	fmt.Printf("t:%d a[%d]:%c b[%d]:%c\n", t, l, a[l], r, b[r])
		//} else if t == 1 {
		//	fmt.Printf("t:%d a[%d]:%c a[%d]:%c\n", t, l, a[l], r, a[r])
		//} else {
		//	fmt.Printf("t:%d b[%d]:%c b[%d]:%c\n", t, l, b[l], r, b[r])
		//}
		if t == 0 && a[l] != b[r] {
			lt, rt, t = l, r, 1
			continue
		} else if t == 1 && a[l] != a[r] {
			l, r, t = lt, rt, 2
			continue
		} else if t == 2 && b[l] != b[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}
