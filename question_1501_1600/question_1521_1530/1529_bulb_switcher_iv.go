package question_1521_1530

// 1529. 灯泡开关 IV
// https://leetcode-cn.com/problems/bulb-switcher-iv/
// Topics: 字符串

func minFlips(target string) int {
	var last, res = false, 0
	for _, c := range target {
		if (c == '1') != last {
			res++
			last = !last
		}
	}
	return res
}
