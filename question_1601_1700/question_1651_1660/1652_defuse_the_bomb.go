package question_1651_1660

// 1652. 拆炸弹
// https://leetcode-cn.com/problems/defuse-the-bomb/
// Topics: 数组

func decrypt(code []int, k int) []int {
	start, end, l, res := 1, k, len(code), make([]int, len(code))

	if k == 0 {
		return res
	} else if k < 0 {
		start, end = l+k, l-1
	}

	for i := start; i <= end; i++ {
		res[0] += code[i]
	}

	for i := 1; i < len(code); i++ {
		res[i] = res[i-1] - code[(start+i-1)%l] + code[(end+i)%l]
	}
	return res
}
