package question_1201_1210

// 1208. 尽可能使字符串相等
// https://leetcode-cn.com/problems/get-equal-substrings-within-budget
// Topics: 数组 None

func abs(a, b uint8) int {
	t := int(a) - int(b)
	if t < 0 {
		return -t
	}
	return t
}

func equalSubstring(s string, t string, maxCost int) int {
	var left, right, max = 0, 0, 0
	for len(s) > right {
		a := abs(s[right], t[right])
		if a > maxCost {
			if left == right {
				left++
				right++
			} else {
				maxCost += abs(s[left], t[left])
				left++
			}
		} else {
			maxCost -= a
			right++
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}
