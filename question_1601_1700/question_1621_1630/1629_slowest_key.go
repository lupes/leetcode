package question_1621_1630

// 1629. 按键持续时间最长的键
// https://leetcode-cn.com/problems/slowest-key/
// Topics: 数组

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var max, res, d = releaseTimes[0], keysPressed[0], 0
	for i := 1; i < len(releaseTimes); i++ {
		d = releaseTimes[i] - releaseTimes[i-1]
		if d > max || (d == max && res < keysPressed[i]) {
			max, res = d, keysPressed[i]
		}
	}
	return res
}
