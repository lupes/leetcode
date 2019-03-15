package question_11_20

var base = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

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
