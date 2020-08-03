package question_791_800

// 796. 旋转字符串
// https://leetcode-cn.com/problems/rotate-string
// Topics:

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A)+1; i++ {
		if A == B {
			return true
		}
		A = A[1:] + A[0:1]
	}
	return false
}
