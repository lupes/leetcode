package question_1511_1520

// 1518. 换酒问题
// https://leetcode-cn.com/problems/water-bottles/
// Topics: 数学 模拟

func numWaterBottles(numBottles int, numExchange int) int {
	var res, tmp = numBottles, numBottles
	for tmp >= numExchange {
		res += tmp / numExchange
		tmp = tmp/numExchange + tmp%numExchange
	}
	return res
}
