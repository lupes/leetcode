package question_561_570

// 565. 数组嵌套
// https://leetcode-cn.com/problems/array-nesting
// Topics: 数组

func arrayNesting(nums []int) int {
	var flag, max = make([]int, len(nums)), 0
	for i := range nums {
		arrayNestingHelper(nums, flag, i)
		if flag[i] > max {
			max = flag[i]
		}
	}
	return max
}

func arrayNestingHelper(nums, flag []int, i int) int {
	switch flag[i] {
	case -1:
		return 0
	case 0:
		flag[i] = -1
		flag[i] = arrayNestingHelper(nums, flag, nums[i]) + 1
	}
	return flag[i]
}
