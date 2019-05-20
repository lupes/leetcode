package question_81_90

func search(nums []int, target int) bool {
	for _, n := range nums {
		if target == n {
			return true
		}
	}
	return false
}
