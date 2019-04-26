package question_51_60

func canJump(nums []int) bool {
	max, size := 0, len(nums)-1
	for i, n := range nums {
		if n+i > max {
			max = n + i
		}
		if n == 0 && max <= i && i != size {
			return false
		}
	}
	return true
}
