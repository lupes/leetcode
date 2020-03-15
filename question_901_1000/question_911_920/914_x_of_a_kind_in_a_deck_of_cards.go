package question_911_920

// 914. 卡牌分组
// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
// Topics: 数组 数学

func hasGroupsSizeX(deck []int) bool {
	var flag = make(map[int]int)
	for _, n := range deck {
		flag[n]++
	}
	var t = 0
	for _, v := range flag {
		if t == 0 {
			t = v
		}
		if t != v && t > v {
			t = gcd(t, v)
		} else if t != v && t < v {
			t = gcd(v, t)
		}
		if t == 1 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if b > 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}
