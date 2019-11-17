package question_41_50

// 46. 全排列
// https://leetcode-cn.com/problems/permutations

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	}
	for i, num := range nums {
		nums[0], nums[i] = nums[i], nums[0]
		for _, t := range permute(nums[1:]) {
			r := append([]int{num}, t...)
			res = append(res, r)
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res
}
