package question_1671_1680

// 1673. 找出最具竞争力的子序列
// https://leetcode-cn.com/problems/find-the-most-competitive-subsequence/
// Topics: 栈 堆 贪心算法 队列

func mostCompetitive(nums []int, k int) []int {
	var stack = make([]int, 0)
	for i, n := range nums {
		if len(stack)+len(nums[i:]) == k {
			return append(stack, nums[i:]...)
		}
		for len(stack) > 0 && n < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			if len(stack)+len(nums[i:]) == k {
				return append(stack, nums[i:]...)
			}
		}
		stack = append(stack, n)
	}
	return stack[:k]
}
