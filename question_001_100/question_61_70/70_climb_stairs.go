package question_61_70

var m = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		r1, ok := m[n-1]
		if !ok {
			r1 = climbStairs(n - 1)
			m[n-1] = r1
		}
		r2, ok := m[n-2]
		if !ok {
			r2 = climbStairs(n - 2)
		}
		return r1 + r2
	}
}
