package question_381_390

// 389. 找不同
// https://leetcode-cn.com/problems/find-the-difference
// Topics: 位运算 哈希表

func findTheDifference(s string, t string) byte {
	var flag = make(map[byte]int)
	for _, c := range s {
		flag[byte(c)]++
	}
	for _, c := range t {
		flag[byte(c)]--
	}
	for k := byte('a'); k <= 'z'; k++ {
		if flag[k] == -1 {
			return k
		}
	}
	return 0
}
