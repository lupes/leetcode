package question_431_440

// 438. 找到字符串中所有字母异位词
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
// Topics: 哈希表

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	var res, l, flag, tmp = make([]int, 0), len(p), make(map[byte]int, 26), make(map[byte]int, 26)
	for i := range p {
		flag[p[i]]++
		tmp[s[i]]++
	}
	tmp[s[l-1]]--
	for i := l - 1; i < len(s); i++ {
		tmp[s[i]]++
		for j := byte('a'); j <= 'z'; j++ {
			if flag[j] != tmp[j] {
				goto Next
			}
		}
		res = append(res, i-l+1)
	Next:
		tmp[s[i-l+1]]--
	}
	return res
}
