package common

// 插入排序
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		value, j := nums[i], i
		for ; j > 0 && value < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = value
	}
	return nums
}
