package question_661_670

// 665. 非递减数列
// https://leetcode-cn.com/problems/non-decreasing-array
// Topics: 数组

func checkPossibility(nums []int) bool {
	var flag, l = 0, len(nums)
	for i := 1; i < l; i++ {
		if nums[i-1] > nums[i] {
			if i == 1 {
				nums[i-1] = nums[i]
			} else if i == l-1 {
				nums[i] = nums[i-1]
			} else {
				if nums[i-1] > nums[i+1] && nums[i] < nums[i-2] {
					return false
				} else if nums[i-1] > nums[i+1] {
					nums[i-1] = nums[i]
				} else {
					nums[i] = nums[i+1]
				}
			}
			flag++
			if flag > 1 {
				return false
			}
		}
	}
	return true
}
