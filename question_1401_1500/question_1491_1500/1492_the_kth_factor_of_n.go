package question_1491_1500

// 1481. 不同整数的最少数目
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
