package question_181_190

// 187. 重复的DNA序列
// https://leetcode-cn.com/problems/repeated-dna-sequences/
// Topics: 位运算 哈希表

func findRepeatedDnaSequences(s string) []string {
	var res []string
	var tmp = make(map[int]bool)
	for i := 0; i < len(s)-9; i++ {
		key := hash(s[i : i+10])
		if b, ok := tmp[key]; ok {
			if b {
				res = append(res, s[i:i+10])
				tmp[key] = false
			}
		} else {
			tmp[key] = true
		}
	}
	return res
}

func hash(s string) int {
	var res = 0
	for _, c := range s {
		switch c {
		case 'A':
			res = res << 2
		case 'C':
			res = res<<2 + 1
		case 'G':
			res = res<<2 + 2
		case 'T':
			res = res<<2 + 3
		}
	}
	return res
}
