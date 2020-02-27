package question_851_860

// 859. 亲密字符串
// https://leetcode-cn.com/problems/buddy-strings
// Topics: 字符串

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	var num = 0
	var a, b byte
	var flag = make([]int, 26)
	var t = false
	for i := range A {
		if A[i] != B[i] {
			num++
			if num == 1 {
				a, b = A[i], B[i]
			} else if num == 2 && (a != B[i] || b != A[i]) {
				return false
			} else if num == 3 {
				return false
			}
		} else {
			flag[A[i]-'a']++
			if flag[A[i]-'a'] >= 2 {
				t = true
			}
		}
	}
	if num == 0 && t {
		return true
	}
	return num == 2
}
