package question_11_20

var base = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

// 17. 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
// Topics: 字符串 回溯算法

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := []string{""}
	for _, c := range digits {
		var tmp []string
		strs := base[c-'2']
		for _, str := range res {
			for _, char := range strs {
				tmp = append(tmp, str+string(char))
			}
		}
		res = tmp
	}
	return res
}
