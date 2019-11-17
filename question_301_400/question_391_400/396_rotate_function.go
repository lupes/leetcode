package question_391_400

// 396. 旋转函数
// https://leetcode-cn.com/problems/rotate-function/
// Topics: 数学

func maxRotateFunction(A []int) int {
	max, tmp, sum, l := 0, 0, 0, len(A)
	for i, n := range A {
		tmp, sum = tmp+i*n, sum+n
	}
	max = tmp
	for i, _ := range A {
		tmp = tmp - l*A[l-i-1] + sum
		if tmp > max {
			max = tmp
		}
	}
	return max
}
