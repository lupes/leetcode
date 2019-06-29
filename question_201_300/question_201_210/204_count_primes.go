package question_201_210

// 204. 计数质数
// https://leetcode-cn.com/problems/count-primes

func countPrimes(n int) int {
	flags := make([]bool, n)
	for i := 2; i < n-1; i++ {
		if !flags[i] {
			for j := i + 1; j < n; j++ {
				if !flags[j] && j%i == 0 {
					flags[j] = true
				}
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !flags[i] {
			count++
		}
	}
	return count
}
