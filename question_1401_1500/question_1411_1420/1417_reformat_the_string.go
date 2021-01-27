package question_1411_1420

// 1417. 重新格式化字符串
// https://leetcode-cn.com/problems/reformat-the-string/
// Topics: 字符串

func reformat(s string) string {
	var numbers, letters []int32
	for _, c := range s {
		if c >= '0' && c <= '9' {
			numbers = append(numbers, c)
		} else {
			letters = append(letters, c)
		}
	}

	if len(numbers) != len(letters) && len(numbers)+1 != len(letters) && len(numbers) != len(letters)+1 {
		return ""
	}

	a, b := numbers, letters
	if len(letters) > len(numbers) {
		b, a = a, b
	}

	var res = make([]byte, 0, len(s))
	for i := 0; i < len(s)/2; i++ {
		res = append(res, byte(a[i]), byte(b[i]))
	}

	if len(s)%2 == 1 {
		res = append(res, byte(a[len(a)-1]))
	}

	return string(res)
}
