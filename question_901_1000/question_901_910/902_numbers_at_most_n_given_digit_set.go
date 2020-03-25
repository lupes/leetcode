package question_901_910

// 902. 最大为 N 的数字组合
// https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set
// Topics: 数学 动态规划

func atMostNGivenDigitSet(D []string, N int) int {
	l, t := 0, N
	for t > 0 {
		l++
		t /= 10
	}
	var nD []int
	for _, d := range D {
		nD = append(nD, int(d[0])-'0')
	}
	return atMostNGivenDigitSetHelper(nD, N, l, 0)
}

func atMostNGivenDigitSetHelper(D []int, N int, i int, t int) int {
	if i == -1 && t != 0 {
		return 1
	} else if i == -1 && t == 0 {
		return 0
	}
	var res = 0
	if t == 0 {
		res += atMostNGivenDigitSetHelper(D, N, i-1, t)
	}
	for _, n := range D {
		x := to(n, i)
		if t+x <= N {
			res += atMostNGivenDigitSetHelper(D, N, i-1, t+x)
		} else {
			break
		}
	}
	return res
}

func to(i, b int) int {
	for j := 0; j < b; j++ {
		i *= 10
	}
	return i
}

func pow(x, y int) int {
	var res = 1
	for j := 0; j < y; j++ {
		res *= x
	}
	return res
}
