package question_31_40

func searchRange(nums []int, target int) []int {
	count := len(nums)
	if count == 0 {
		return []int{-1, -1}
	}
	l, r := 0, count
	for {
		c := (l + r) / 2
		if nums[c] == target {

		}
		if nums[c] > target {
			r = c
		}
	}
}
