package question_491_500

// 503. 下一个更大元素 II
// https://leetcode-cn.com/problems/next-greater-element-ii/
// Topics: 栈

func nextGreaterElements(nums []int) []int {
	var res, l, flag = make([]int, len(nums)), len(nums), false
	for i, n := range nums {
		res[i], flag = -1, false
		for j := l - 1; j >= 0; j-- {
			if nums[j] > n {
				res[i], flag = nums[j], true
			}
			if i == j && flag {
				break
			}
		}
	}
	return res
}
