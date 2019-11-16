package question_41_50

// 43. 字符串相乘
// https://leetcode-cn.com/problems/multiply-strings

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	size1, size2 := len(num1), len(num2)
	m, t := uint8(0), uint8(0)
	index := 0
	size3 := size1 + size2 + 1
	var tmp = make([]uint8, size3)
	for i := size2 - 1; i >= 0; i-- {
		for j := size1 - 1; j >= 0; j-- {
			m = uint8(num2[i]-'0') * uint8(num1[j]-'0')
			index = size3 - i - j - 3
			t = tmp[index] + m
			tmp[index] = t % 10
			tmp[index+1] += t / 10
		}
	}

	result := ""
	for i := 0; i < len(tmp)-1; i++ {
		result = string(tmp[i]%10+'0') + result
		tmp[i+1] += tmp[i] / 10
	}
	for i, c := range result {
		if c != '0' {
			return result[i:]
		}
	}
	return result
}
