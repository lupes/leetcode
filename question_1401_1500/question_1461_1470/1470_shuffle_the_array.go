package question_1461_1470

// 1470. 重新排列数组
// https://leetcode-cn.com/problems/shuffle-the-array/
// Topics: 数组

func shuffle(nums []int, n int) []int {
	var res = make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[n+i])
	}
	return res
}
