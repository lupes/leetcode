package question_1371_1380

// 1395. 统计作战单位数
// https://leetcode-cn.com/problems/count-number-of-teams/
// Topics: 栈 设计

func numTeams(rating []int) int {
	var res int
	for i := range rating {
		var leftUp, leftDown = 0, 0
		for j := 0; j < i; j++ {
			if rating[i] > rating[j] {
				leftDown++
			}
			if rating[i] < rating[j] {
				leftUp++
			}
		}

		var rightUp, rightDown = 0, 0
		for j := i + 1; j < len(rating); j++ {
			if rating[i] > rating[j] {
				rightDown++
			}
			if rating[i] < rating[j] {
				rightUp++
			}
		}
		res += leftUp*rightDown + leftDown*rightUp
	}
	return res
}
