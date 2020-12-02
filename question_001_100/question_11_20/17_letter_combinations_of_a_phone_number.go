package question_11_20

// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
// Topics: 字符串 回溯算法

var base = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	var res []string
	letterCombinationsHelper(digits, &res, "")
	return res
}

func letterCombinationsHelper(digits string, res *[]string, tmp string) {
	if digits == "" {
		*res = append(*res, tmp)
		return
	}
	chars := base[digits[0]-'2']
	for _, c := range chars {
		letterCombinationsHelper(digits[1:], res, tmp+string(c))
	}
	return
}
