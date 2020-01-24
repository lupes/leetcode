package question_721_730

// 728. 自除数
// https://leetcode-cn.com/problems/self-dividing-numbers
// Topics: 数学

func selfDividingNumbers(left int, right int) []int {
	var res []int
Next:
	for i := left; i <= right; i++ {
		tmp := i
		for tmp > 0 {
			if tmp%10 == 0 || i%(tmp%10) != 0 {
				continue Next
			}
			tmp /= 10
		}
		res = append(res, i)
	}
	return res
}
