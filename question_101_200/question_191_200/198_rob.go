package question_191_200

// 198. 打家劫舍
// https://leetcode-cn.com/problems/house-robber/

func rob(nums []int) int {
	var last, now, tmp = 0, 0, 0
	for _, n := range nums {
		tmp = now
		if last+n > now {
			now = last + n
		}
		last = tmp
	}
	return now
}
