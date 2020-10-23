package question_1331_1340

// 1342. 将数字变成 0 的操作次数
// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
// Topics: 位运算

func numberOfSteps(num int) int {
	var res = 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		res++
	}
	return res
}
