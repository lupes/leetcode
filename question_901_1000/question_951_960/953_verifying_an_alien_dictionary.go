package question_951_960

// 953. 验证外星语词典
// https://leetcode-cn.com/problems/verifying-an-alien-dictionary
// Topics: 哈希表

func trans(word string, flag map[int32]byte) string {
	var res = make([]byte, len(word))
	for i, c := range word {
		res[i] = flag[c]
	}
	return string(res)
}

func isAlienSorted(words []string, order string) bool {
	var flag = make(map[int32]byte)
	for i, c := range order {
		flag[c] = byte('a' + i)
	}

	a, b := trans(words[0], flag), ""
	for i := 1; i < len(words); i++ {
		b = trans(words[i], flag)
		if a > b {
			return false
		}
		a = b
	}
	return true
}
