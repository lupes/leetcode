package question_431_440

// 438. 找到字符串中所有字母异位词
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
// Topics: 哈希表

func findAnagrams(s string, p string) []int {
	var res, ls, lp = make([]int, 0), len(s), len(p)
	if lp > ls {
		return nil
	}
	var sFlag, pFlag = make([]int, 'z'+1), make([]int, 'z'+1)
	for i := range p {
		pFlag[p[i]]++
		sFlag[s[i]]++
	}
	sFlag[s[lp-1]]--
	for i := lp - 1; i < ls; i++ {
		sFlag[s[i]]++
		for j := 'a'; j <= 'z'; j++ {
			if pFlag[j] != sFlag[j] {
				goto Next
			}
		}
		res = append(res, i-lp+1)
	Next:
		sFlag[s[i-lp+1]]--
	}
	return res
}
