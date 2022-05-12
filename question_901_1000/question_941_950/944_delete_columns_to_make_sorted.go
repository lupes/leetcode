package question_941_950

// 944. 删列造序
// https://leetcode-cn.com/problems/delete-columns-to-make-sorted
// Topics: 贪心算法

func minDeletionSize(A []string) int {
	l, n := len(A[0]), 0
Next:
	for i := 0; i < l; i++ {
		t := A[0][i]
		for _, a := range A {
			if a[i] < t {
				n++
				continue Next
			}
			t = a[i]
		}
	}
	return n
}
