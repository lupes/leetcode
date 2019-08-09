package question_311_320

// 319. 灯泡开关
// https://leetcode-cn.com/problems/bulb-switcher/

func bulbSwitch(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if i*i <= n {
			res++
		} else {
			break
		}
	}
	return res
}
