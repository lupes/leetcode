package question_441_450

// 442. 数组中重复的数据
// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/

func findDuplicates(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		} else if nums[i] > 0 && i != nums[i]-1 && nums[i] == nums[nums[i]-1] {
			res = append(res, nums[i])
			nums[i] = 0
		}
	}
	return res
}
