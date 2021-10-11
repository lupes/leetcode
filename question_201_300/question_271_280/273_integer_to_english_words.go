package question_271_280

import (
	"strings"
)

// 273. 整数转换英文表示
// https://leetcode-cn.com/problems/integer-to-english-words
// Topics: 数学 字符串

var low = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
var mid = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var high = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var flag = []string{"", "Thousand", "Million", "Billion"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var res = ""
	for i := 0; num > 0; i++ {
		tmp := num % 1000
		if tmp > 0 {
			res = numberToWordsHelper(tmp) + " " + flag[i] + " " + res
		}

		num /= 1000
	}
	return strings.TrimSpace(res)
}

func numberToWordsHelper(num int) string {
	var res string
	if num > 99 {
		res += low[num/100] + " Hundred "
	}
	num %= 100

	if num > 19 {
		res += high[num/10] + " "
		num %= 10
	}
	if num > 9 {
		res += mid[num-10]
		num = 0
	}

	res += low[num]

	return strings.TrimSpace(res)
}
