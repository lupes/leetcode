package question_781_790

// 784. 字母大小写全排列
// https://leetcode-cn.com/problems/letter-case-permutation
// Topics: 位运算 回溯算法

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{""}
	}
	var res []string
	for i, c := range S {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			tmp := letterCasePermutation(S[i+1:])
			for _, t := range tmp {
				if 'a' <= c && c <= 'z' {
					res = append(res, S[:i]+string(c-'a'+'A')+t)
				} else {
					res = append(res, S[:i]+string(c-'A'+'a')+t)
				}
				res = append(res, S[:i+1]+t)
			}
			return res
		}
	}
	return []string{S}
}
