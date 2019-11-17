package question_181_190

// 187. 重复的DNA序列
// https://leetcode-cn.com/problems/repeated-dna-sequences/
// Topics: 位运算 哈希表

func findRepeatedDnaSequences(s string) []string {
	var res []string
	var tmp = make(map[string]bool)
	var key string
	for i := 0; i < len(s)-9; i++ {
		key = s[i : i+10]
		if b, ok := tmp[key]; ok {
			if b {
				res = append(res, key)
				tmp[key] = false
			}
		} else {
			tmp[key] = true
		}
	}
	return res
}
