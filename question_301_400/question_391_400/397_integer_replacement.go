package question_391_400

// 397. 整数替换
// https://leetcode-cn.com/problems/integer-replacement/

func integerReplacement(n int) int {
	res := 0
	for ; n > 1; res++ {
		if n == 3 {
			n++
		}
		if n%2 == 0 {
			n /= 2
		} else if ((n+1)/2)%2 == 0 {
			n += 1
		} else {
			n -= 1
		}
	}
	return res
}
