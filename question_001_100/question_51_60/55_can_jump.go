package question_51_60

func canJump(nums []int) bool {
	max := 0
	for i, n := range nums {
		if max >= i {
			if n+i > max {
				max = n + i
			}
		} else {
			return false
		}
	}
	return true
}
