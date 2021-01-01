package question_1211_1220

// 1217. 玩筹码
// https://leetcode-cn.com/problems/play-with-chips
// Topics: 贪心算法 数组 数学

func minCostToMoveChips(chips []int) int {
	var odd, even, min = 0, 0, 0
	for _, n := range chips {
		if n%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	min = odd
	if even < min {
		min = even
	}
	return min
}
