package question_1001_1010

// 1007. 行相等的最少多米诺旋转
// https://leetcode-cn.com/problems/minimum-domino-rotations-for-equal-row
// Topics: 贪心算法 数组

func minDominoRotations(A []int, B []int) int {
	r := min(minDominoRotationsHelper(A, B), len(A)+1)
	r = min(minDominoRotationsHelper(B, A), r)

	A[0], B[0] = B[0], A[0]
	r = min(minDominoRotationsHelper(A, B)+1, r)
	r = min(minDominoRotationsHelper(B, A)+1, r)

	if r == len(A)+1 {
		return -1
	}

	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDominoRotationsHelper(A []int, B []int) int {
	t, r := A[0], 0
	for i, n := range A {
		if n != t && B[i] == t {
			r++
		} else if n != t && B[i] != t {
			return len(A) + 1
		}
	}
	return r
}
