package question_471_480

// 473. 火柴拼正方形
// https://leetcode-cn.com/problems/matchsticks-to-square/
// Topics: 深度优先搜索

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}
	r := 0
	for _, n := range nums {
		r += n
	}
	if r%4 != 0 {
		return false
	}
	return makeSquareHelp(nums, r/4, 0, 0, 0, 0, 0)
}

func makeSquareHelp(nums []int, a, a1, a2, a3, a4, i int) bool {
	if len(nums) == i && a1 == a && a2 == a && a3 == a && a4 == a {
		return true
	}
	if a1 > a || a2 > a || a3 > a || a4 > a {
		return false
	}
	return makeSquareHelp(nums, a, a1+nums[i], a2, a3, a4, i+1) ||
		makeSquareHelp(nums, a, a1, a2+nums[i], a3, a4, i+1) ||
		makeSquareHelp(nums, a, a1, a2, a3+nums[i], a4, i+1) ||
		makeSquareHelp(nums, a, a1, a2, a3, a4+nums[i], i+1)
}
