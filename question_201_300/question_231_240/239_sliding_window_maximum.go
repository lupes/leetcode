package question_231_240

// 239. 滑动窗口最大值
// https://leetcode-cn.com/problems/sliding-window-maximum
// Topics: 队列 数组 滑动窗口 单调队列 堆

func maxSlidingWindow(nums []int, k int) []int {
	var stack = make([]int, 0, len(nums)-k)
	var res = make([]int, 0, len(nums)-k)
	for i, n := range nums {
		for len(stack) > 0 && n >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

		if stack[0] == i-k {
			stack = stack[1:]
		}

		if i >= k-1 {
			res = append(res, nums[stack[0]])
		}
	}
	return res
}
