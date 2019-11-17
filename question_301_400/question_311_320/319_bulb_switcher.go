package question_311_320

import "math"

// 319. 灯泡开关
// https://leetcode-cn.com/problems/bulb-switcher/
// Topics: 脑筋急转弯 数学

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
