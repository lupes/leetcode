package question_151_160

func findMin(nums []int) int {
	size := len(nums)
	l, r := 0, size-1
	for r > l {
		if nums[r] > nums[l] {
			return nums[l]
		} else if l+1 == r {
			return nums[r]
		} else {
			c := (l + r) / 2
			if nums[c] > nums[l] {
				l = c
			} else {
				r = c
			}
		}
	}
	return nums[0]
}
