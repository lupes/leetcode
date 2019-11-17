package question_281_290

// 287. 寻找重复数
// https://leetcode-cn.com/problems/find-the-duplicate-number/
// Topics: 数组 双指针 二分查找

func findDuplicate(nums []int) int {
	for i, n := range nums {
		for _, m := range nums[i+1:] {
			if n == m {
				return n
			}
		}
	}
	return 0
}
