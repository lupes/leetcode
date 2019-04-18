package question_41_50

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	}

Continue:
	for i, num := range nums {
		// 去重当某一个数字和前面数字相同时不交换
		for _, t := range nums[:i] {
			if t == num {
				continue Continue
			}
		}
		nums[0], nums[i] = nums[i], nums[0]
		if len(nums[1:]) == 1 {
			res = append(res, []int{num, nums[1]})
		} else {
			for _, t := range permuteUnique(nums[1:]) {
				r := append([]int{num}, t...)
				res = append(res, r)
			}
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res
}
