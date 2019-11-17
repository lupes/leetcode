package question_201_210

// 204. 计数质数
// https://leetcode-cn.com/problems/count-primes
// Topics: 哈希表 数学

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	flags := make([]bool, n)
	count := n / 2
	for i := 3; i*i < n; i += 2 {
		if !flags[i] {
			for j := i * i; j < n; j += 2 * i {
				if !flags[j] && j%i == 0 {
					count--
					flags[j] = true
				}
			}
		}
	}
	return count
}
