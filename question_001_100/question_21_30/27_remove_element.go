package question_0011_0020

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v == val {
			continue
		}
		nums[index] = v
		index++
	}
	return index
}
