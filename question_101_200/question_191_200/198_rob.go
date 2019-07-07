package question_191_200

// 198. 打家劫舍
// https://leetcode-cn.com/problems/house-robber/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var data = make(map[int]int)
	return robDfs(data, nums)
}

func robDfs(data map[int]int, nums []int) int {
	l := len(nums)
	if r, ok := data[l]; ok {
		return r
	}
	var res int
	if l == 1 {
		res = nums[0]
	} else if l == 2 {
		res = nums[0]
		if nums[1] > nums[0] {
			res = nums[1]
		}
	} else {
		res = rob(nums[1:])
		s2 := nums[0] + rob(nums[2:])
		if s2 > res {
			res = s2
		}
	}
	data[l] = res
	return res
}
