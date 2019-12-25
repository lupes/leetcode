package question_1051_1060

// 1052. 爱生气的书店老板
// https://leetcode-cn.com/problems/grumpy-bookstore-owner
// Topics: 数组 None

func maxSatisfied(customers []int, grumpy []int, X int) int {
	var left, right, sum, now, max, gru = 0, 0, 0, 0, 0, 0
	for right < len(customers) {
		if right-left < X {
			if grumpy[right] == 1 {
				now += customers[right]
				gru += customers[right]
			}
			sum += customers[right]
			right++
		} else {
			if grumpy[left] == 1 {
				now -= customers[left]
			}
			left++
		}
		if now > max {
			max = now
		}
	}
	return sum - gru + max
}
