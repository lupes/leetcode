package question_741_750

// 748. 最短完整词
// https://leetcode-cn.com/problems/shortest-completing-word
// Topics: 哈希表

func shortestCompletingWord(licensePlate string, words []string) string {
	var src = [26]int{}
	for _, c := range licensePlate {
		if c >= 'a' && c <= 'z' {
			src[c-'a']++
		} else if c >= 'A' && c <= 'Z' {
			src[c-'A']++
		}
	}

	var flag, res = [26]int{}, "12345678901234567890"
	for _, word := range words {
		for _, c := range word {
			if c >= 'a' && c <= 'z' {
				flag[c-'a']++
			} else if c >= 'A' && c <= 'Z' {
				flag[c-'A']++
			}
		}
		f := true
		for i := 0; i < 26; i++ {
			if flag[i] < src[i] {
				f = false
			}
			flag[i] = 0
		}
		if f && len(res) > len(word) {
			res = word
		}
	}
	return res
}
