package question_281_290

// 283. 移动零
// https://leetcode-cn.com/problems/move-zeroes/
// Topics: 数组 双指针

func moveZeroes(nums []int) {
	zeroIndex := 0
	for i, n := range nums {
		if n != 0 {
			if zeroIndex != i {
				nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			}
			zeroIndex++
		}
	}
}
