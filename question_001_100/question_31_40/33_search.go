package question_31_40

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	if nums[0] > target {
		for i := size - 1; i > 0; i-- {
			if nums[i] == target {
				return i
			}
			if i < size-1 {
				if nums[i] > nums[i+1] {
					return -1
				}
			}
		}
	} else {
		for i := 0; i < size; i++ {
			if nums[i] == target {
				return i
			}
			if i > 0 {
				if nums[i] < nums[i-1] {
					return -1
				}
			}
		}
	}
	return -1
}
