package question_181_190

// 189. 旋转数组
// https://leetcode-cn.com/problems/rotate-array/
// Topics: 数组

func rotate(nums []int, k int) {
	l := len(nums)
	n := k % l
	if n == 0 {
		return
	} else {
		for i, n := range append(nums[l-n:], nums[:l-n]...) {
			nums[i] = n
		}
	}
}
