package question_391_400

// 397. 整数替换
// https://leetcode-cn.com/problems/integer-replacement/

var m = map[int]int{1: 0}

func integerReplacement(n int) int {
	if res, ok := m[n]; ok {
		return res
	}
	res, t := 0, 0
	if n%2 == 0 {
		res = integerReplacement(n / 2)
	} else {
		res, t = integerReplacement(n+1), integerReplacement(n-1)
		if res > t {
			res = t
		}
	}
	return res + 1
}
