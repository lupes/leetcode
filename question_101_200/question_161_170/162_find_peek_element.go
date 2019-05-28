package question_161_170

func findPeakElement(nums []int) int {
	size := len(nums)
	for i, _ := range nums {
		if (size == 1) ||
			(i == 0 && (size > 1 && nums[1] < nums[0])) ||
			(i == size-1 && i > 0 && nums[i-1] < nums[i]) ||
			(i < size-1 && i > 0 && nums[i] > nums[i-1] && nums[i] > nums[i+1]) {
			return i
		}
	}
	return -1
}

// (i < size -1 && nums[i] > nums[i+1]) || (i == size -1) {
//			return i
//		}
