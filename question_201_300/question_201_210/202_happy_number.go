package question_201_210

// 202. 快乐数
// https://leetcode-cn.com/problems/happy-number/
// Topics: 哈希表 数学

var m = map[int]int{
	0: 0,
	1: 1,
	2: 4,
	3: 9,
	4: 16,
	5: 25,
	6: 36,
	7: 49,
	8: 64,
	9: 81,
}

func isHappy(n int) bool {
	var ms = make(map[int]struct{})
	for n != 1 {
		var t int
		for n > 0 {
			t += m[n%10]
			n /= 10
		}
		n = t
		_, ok := ms[t]
		if ok {
			return false
		}
		ms[n] = struct{}{}
	}
	return true
}
