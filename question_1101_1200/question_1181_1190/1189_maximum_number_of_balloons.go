package question_1181_1190

// 1189. “气球” 的最大数量
// https://leetcode-cn.com/problems/maximum-number-of-balloons
// Topics: 哈希表 字符串

func maxNumberOfBalloons(text string) int {
	var flag = make(map[int32]int)
	for _, c := range text {
		flag[c]++
	}
	var min = len(text)
	for _, c := range "balon" {
		if c == 'l' || c == 'o' {
			flag[c] /= 2
		}
		if flag[c] < min {
			min = flag[c]
		}
	}
	return min
}
