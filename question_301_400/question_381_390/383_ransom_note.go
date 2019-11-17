package question_381_390

// 383. 赎金信
// https://leetcode-cn.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	var m = make(map[int32]int)
	for _, u := range magazine {
		m[u]++
	}
	for _, u := range ransomNote {
		if i, ok := m[u]; !ok || i == 0 {
			return false
		} else {
			m[u]--
		}
	}
	return true
}
