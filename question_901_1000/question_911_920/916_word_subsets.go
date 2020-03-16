package question_911_920

// 916. 单词子集
// https://leetcode-cn.com/problems/word-subsets
// Topics: 字符串

func wordSubsets(A []string, B []string) []string {
	var res []string
	var BFlag = make(map[int32]int)
	for i := 0; i < len(B); i++ {
		var f = make(map[int32]int)
		for _, c := range B[i] {
			f[c]++
			if f[c] > BFlag[c] {
				BFlag[c] = f[c]
			}
		}
	}
Next:
	for _, a := range A {
		var flag = make(map[int32]int)
		for _, c := range a {
			flag[c]++
		}
		for k, v := range BFlag {
			if flag[k] < v {
				continue Next
			}
		}
		res = append(res, a)
	}
	return res
}
