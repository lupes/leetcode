package question_1331_1340

// 1344. 时钟指针的夹角
// https://leetcode-cn.com/problems/angle-between-hands-of-a-clock
// Topics: 数学

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	mAngle := float64(minutes * 6)
	hAngle := float64(hour*30) + (float64(minutes)/60)*30
	var res = float64(0)
	if mAngle == hAngle {
		return res
	} else if mAngle > hAngle {
		res = mAngle - hAngle
	} else {
		res = hAngle - mAngle
	}
	if res > 180 {
		res = 360 - res
	}
	return res
}
