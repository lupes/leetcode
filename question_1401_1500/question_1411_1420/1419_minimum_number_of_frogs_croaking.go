package question_1411_1420

// 1419. 数青蛙
// https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/
// Topics: 字符串

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}

	var flag = map[string]int{"": len(croakOfFrogs), "c": 0, "cr": 0, "cro": 0, "croa": 0, "croak": 0}
	var m = map[int32]string{'c': "", 'r': "c", 'o': "cr", 'a': "cro", 'k': "croa"}

	var res, tmp = 0, 0
	for _, c := range croakOfFrogs {
		if flag[m[c]] == 0 {
			return -1
		}

		flag[m[c]]--
		flag[m[c]+string(c)]++

		if c == 'c' {
			tmp += 1
			if tmp > res {
				res = tmp
			}
		} else if c == 'k' {
			tmp--
		}
	}

	for k, n := range flag {
		if k != "" && k != "croak" && n > 0 {
			return -1
		}
	}

	return res
}
