package question_1361_1370

// 1370. 上升下降字符串
// https://leetcode-cn.com/problems/increasing-decreasing-string/
// Topics: 排序 字符串

func sortString(s string) string {
	var flag = make(map[byte]int, 26)
	for _, c := range s {
		flag[byte(c)]++
	}
	var res = make([]byte, 0, len(s))
	for len(res) < len(s) {
		for c := byte('a'); c <= 'z'; c++ {
			n := flag[c]
			if n > 0 {
				res = append(res, c)
				flag[c]--
			}
		}
		for c := byte('z'); c >= 'a'; c-- {
			n := flag[c]
			if n > 0 {
				res = append(res, c)
				flag[c]--
			}
		}
	}
	return string(res)
}
