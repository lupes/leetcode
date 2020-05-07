package question_1031_1040

// 1033. 移动石子直到连续
// https://leetcode-cn.com/problems/moving-stones-until-consecutive
// Topics: 脑筋急转弯

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	if a+1 == b && b+1 == c {
		return []int{0, 0}
	} else if a+1 == b {
		return []int{1, c - b - 1}
	} else if b+1 == c {
		return []int{1, b - a - 1}
	} else if a+2 == b {
		return []int{1, c - b}
	} else if b+2 == c {
		return []int{1, b - a}
	} else {
		return []int{2, c - b - 1 + b - a - 1}
	}
}
