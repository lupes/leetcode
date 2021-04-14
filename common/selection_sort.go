package common

// 选择排序
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		min, index := nums[i], i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min, index = nums[j], j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}
