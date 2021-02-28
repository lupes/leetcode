package question_1491_1500

// 1492. n 的第 k 个因子
// https://leetcode-cn.com/problems/the-kth-factor-of-n/
// Topics: 数学

func kthFactor(n int, k int) int {
	var num = 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			num += 1
			if num == k {
				return i
			}
		}
	}
	return -1
}
