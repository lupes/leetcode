package question_1321_1330

// 1323. 6 和 9 组成的最大数字
// https://leetcode-cn.com/problems/maximum-69-number/
// Topics: 数学

func maximum69Number(num int) int {
	var bs []int
	for num > 0 {
		bs = append(bs, num%10)
		num /= 10
	}

	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == 6 {
			bs[i] = 9
			break
		}
	}

	var res = 0
	for i := len(bs) - 1; i >= 0; i-- {
		res = res*10 + bs[i]
	}
	return res
}
