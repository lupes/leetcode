package offer

// 剑指 Offer II 079. 所有子集
// https://leetcode-cn.com/problems/TVdhkn/
// Topics: 位运算 数组 回溯

func subsets(nums []int) [][]int {
	var l = len(nums)
	var res [][]int
	for i := 0; i < 1<<l; i++ {
		var tmp []int
		for j, t := 0, i; j < l; j, t = j+1, t>>1 {
			if t&1 == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		res = append(res, tmp)
	}

	return res
}

func subsetsHelper(nums []int, i int) [][]int {
	if i >= len(nums) {
		return [][]int{{}}
	}
	var res [][]int
	tmp := subsetsHelper(nums, i+1)
	for _, t := range tmp {
		res = append(res, t)
		res = append(res, append([]int{nums[i]}, t...))
	}
	return res
}
