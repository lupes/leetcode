package question_1621_1630

// 1629. 按键持续时间最长的键
// https://leetcode-cn.com/problems/slowest-key/
// Topics: 数组

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var flag [26]int
	var d = 0
	flag[keysPressed[0]-'a'] = releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		d = releaseTimes[i] - releaseTimes[i-1]
		if d > flag[keysPressed[i]-'a'] {
			flag[keysPressed[i]-'a'] = d
		}
	}
	var max, res = 0, byte(0)
	for i, num := range flag {
		if num >= max {
			res, max = byte(i)+'a', num
		}
	}
	return res
}
