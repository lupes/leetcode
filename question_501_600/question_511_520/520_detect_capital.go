package qustion_511_520

// 520. 检测大写字母
// https://leetcode-cn.com/problems/detect-capital/
// Topics: 字符串

func detectCapitalUse(word string) bool {
	num, flag, l := 0, false, len(word)
	for i, c := range word {
		if 'A' <= c && c <= 'Z' {
			if i == 0 {
				flag = true
			}
			num++
		}
	}
	switch {
	case num == 0:
		return true
	case num == 1 && flag == true:
		return true
	case num == l:
		return true
	}
	return false
}
