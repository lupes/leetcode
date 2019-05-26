package question_151_160

func findMin(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	var min = nums[0]
	for _, n := range nums {
		if min > n {
			min = n
		}
	}

	//	l, r := 0, size-1
	//	for r > l {
	//		if nums[r] > nums[l] {
	//			return nums[l]
	//		} else {
	//			c := (l + r) / 2
	//			if nums[c] > nums[l] {
	//				l = c
	//			} else {
	//				r = c
	//			}
	//		}
	//	}
	return min
}
