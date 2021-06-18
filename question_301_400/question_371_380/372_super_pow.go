package question_371_380

// 372. 超级次方
// https://leetcode-cn.com/problems/super-pow
// Topics: 数学

func superPow(a int, b []int) int {
	var mod, l, last, next, res = 1337, len(b), a % 1337, 1, 1

	for j := 0; j < b[l-1]; j++ {
		res = res * last % mod
	}

	for i := l - 2; i >= 0; i-- {
		for j := 0; j < 10; j++ {
			next = next * last % mod
		}
		for j := 0; j < b[i]; j++ {
			res = res * next % mod
		}
		last, next = next, 1
	}

	return res
}
