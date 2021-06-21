package question_461_470

import (
	"math/rand"
	"time"
)

// 470. 用 Rand7() 实现 Rand10()
// https://leetcode-cn.com/problems/implement-rand10-using-rand7
// Topics: 随机 拒绝采样

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7)
}

func rand10() int {
	a, b := rand7(), rand7()
	if a > 4 && b < 4 {
		return rand10()
	}
	return (a+b)%10 + 1
}
