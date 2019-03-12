package question_11_20

func romanToInt(s string) int {
	i := 0
	res := 0
	size := len(s)
	for i < size {
		if s[i] == 'M' {
			res += 1000
			i++
		} else if s[i] == 'D' {
			res += 500
			i++
		} else if s[i] == 'C' {
			res, i = helper2(s, i, res, size, 'M', 'D', 100)
		} else if s[i] == 'L' {
			res += 50
			i++
		} else if s[i] == 'X' {
			res, i = helper2(s, i, res, size, 'C', 'L', 10)
		} else if s[i] == 'V' {
			res += 5
			i++
		} else {
			res, i = helper2(s, i, res, size, 'X', 'V', 1)
		}
	}
	return res
}

func helper2(s string, i, res, size int, c1, c2 uint8, num int) (int, int) {
	if size > i+1 {
		if s[i+1] == c1 {
			return res + 9*num, i + 2
		} else if s[i+1] == c2 {
			return res + 4*num, i + 2
		} else {
			return res + num, i + 1
		}
	} else {
		return res + num, i + 1
	}
}
