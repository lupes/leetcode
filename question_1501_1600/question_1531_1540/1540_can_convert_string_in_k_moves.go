package question_1531_1540

// 1540. K 次操作转变字符串
// https://leetcode-cn.com/problems/can-convert-string-in-k-moves/
// Topics: 贪心算法 字符串

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}

	var kFlag = [27]int{len(s)}
	for i := 1; i < 27; i++ {
		kFlag[i] = k / 26
		if k%26 >= i {
			kFlag[i]++
		}
	}

	for i := 0; i < len(s); i++ {
		n := (t[i] - s[i] + 26) % 26
		kFlag[n]--
		if kFlag[n] < 0 {
			return false
		}
	}
	return true
}
