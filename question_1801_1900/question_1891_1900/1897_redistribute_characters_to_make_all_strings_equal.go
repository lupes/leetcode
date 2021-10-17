package question_1891_1900

// 1897. 重新分配字符使所有字符串都相等
// https://leetcode-cn.com/problems/redistribute-characters-to-make-all-strings-equal/
// Topics: 哈希表 字符串 计数

func makeEqual(words []string) bool {
	var flag = make([]int, 'z'+1)
	for _, word := range words {
		for _, c := range word {
			flag[c]++
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if flag[i]%len(words) != 0 {
			return false
		}
	}
	return true
}
