package question_0011_0020

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	index := 0
	for i, v := range nums {
		if i == 0 || v == nums[index] {
			continue
		}
		index++
		if i != index {
			nums[index] = v
		}
	}
	return index + 1
}
