package question_651_660

// 657. 机器人能否返回原点
// https://leetcode-cn.com/problems/robot-return-to-origin
// Topics: 字符串

func judgeCircle(moves string) bool {
	var m = make(map[int32]int)
	for _, c := range moves {
		m[c]++
	}
	if m['U'] == m['D'] && m['L'] == m['R'] {
		return true
	}
	return false
}
