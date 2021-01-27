package question_1411_1420

// 1419. 数青蛙
// https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/
// Topics: 字符串

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}

	var c, r, o, a, k, res, tmp int
	for _, char := range croakOfFrogs {
		switch char {
		case 'c':
			c, tmp = c+1, tmp+1
			if tmp > res {
				res = tmp
			}
		case 'r':
			r++
		case 'o':
			o++
		case 'a':
			a++
		case 'k':
			k, tmp = k+1, tmp-1
		}

		if !(c >= r && r >= o && o >= a && a >= k) {
			return -1
		}
	}

	if c != r || r != o || o != a || a != k {
		return -1
	}

	return res
}
