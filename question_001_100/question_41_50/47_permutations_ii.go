package question_41_50

// 47. 全排列 II
// https://leetcode-cn.com/problems/permutations-ii

func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return [][]int{nums}
	}

Continue:
	for i, num := range nums {
		// 去重当某一个数字和前面数字相同时不交换
		for _, t := range nums[:i] {
			if t == num {
				continue Continue
			}
		}
		nums[0], nums[i] = nums[i], nums[0]
		for _, t := range permuteUnique(nums[1:]) {
			res = append(res, append(t, num))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return res
}
