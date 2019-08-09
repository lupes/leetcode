package question_311_320

import "math"

// 319. 灯泡开关
// https://leetcode-cn.com/problems/bulb-switcher/

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
