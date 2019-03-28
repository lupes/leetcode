package question_31_40

func searchRange(nums []int, target int) []int {
	count := len(nums)
	res := []int{-1, -1}
	if count == 0 || target < nums[0] || target > nums[count-1] {
		return res
	}
	l, r := 0, count-1
	for l <= r {
		c := (l + r) / 2
		if nums[c] == target {
			for i := 0; c-i > -1 && nums[c-i] == target; i++ {
				res[0] = c - i
			}
			for i := 0; c+i < count && nums[c+i] == target; i++ {
				res[1] = c + i
			}
			break
		} else if nums[c] > target {
			r = c - 1
		} else {
			l = c + 1
		}
	}
	return res
}
