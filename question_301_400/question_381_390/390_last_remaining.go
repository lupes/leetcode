package question_381_390

// 390. 消除游戏
// https://leetcode-cn.com/problems/elimination-game/

func lastRemaining(n int) int {
	var a = make([]int, n)
	for i, _ := range a {
		a[i] = i + 1
	}
	for i := 1; len(a) > 1; i++ {
		if i%2 == 1 {
			a = leftRemove(a)
		} else {
			a = rightRemove(a)
		}
	}
	return a[0]
}

func leftRemove(a []int) []int {
	l, r := 0, 1
	for ; r < len(a); l, r = l+1, r+2 {
		a[l] = a[r]
	}
	return a[:l]
}

func rightRemove(a []int) []int {
	l, r := len(a)-2, len(a)-1
	for ; l >= 0; l, r = l-2, r-1 {
		a[r] = a[l]
	}
	return a[r+1:]
}
