package question_211_220

// 219. 存在重复元素 II
// https://leetcode-cn.com/problems/contains-duplicate-ii/
// Topics: 数组 哈希表

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= k && i+j < len(nums); j++ {
			if nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}
