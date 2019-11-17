package question_291_300

import "fmt"

// 299. 猜数字游戏
// https://leetcode-cn.com/problems/bulls-and-cows/
// Topics: 哈希表

func getHint(secret string, guess string) string {
	a, b := 0, 0
	sm, gm := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			a++
		} else {
			sm[secret[i]]++
			gm[guess[i]]++
		}
	}
	for i := uint8('0'); i <= '9'; i++ {
		if sm[i] > gm[i] {
			b += gm[i]
		} else {
			b += sm[i]
		}
	}
	return fmt.Sprintf("%dA%dB", a, b)
}
