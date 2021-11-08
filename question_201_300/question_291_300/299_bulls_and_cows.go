package question_291_300

import "fmt"

// 299. 猜数字游戏
// https://leetcode-cn.com/problems/bulls-and-cows/
// Topics: 哈希表 字符串 计数

func getHint(secret string, guess string) string {
	a, b := 0, 0
	var flag [10][2]int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			flag[secret[i]-'0'][0]++
			flag[guess[i]-'0'][1]++
		}
	}
	for i := 0; i <= 9; i++ {
		if flag[i][0] > flag[i][1] {
			b += flag[i][1]
		} else {
			b += flag[i][0]
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
