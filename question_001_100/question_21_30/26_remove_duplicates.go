package question_0011_0020

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	offset := 0
	for i := 1; i < size; i++ {
		if nums[i] == nums[i-1-offset] {
			offset += 1
		} else {
			nums[i-offset] = nums[i]
		}
	}
	return size - offset
}
