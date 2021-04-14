package common

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] < nums[0] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[0], nums[i] = nums[i], nums[0]
	quickSort(nums[:i])
	quickSort(nums[i+1:])
	return nums
}
