package question_211_220

// 219. 存在重复元素 II
// https://leetcode-cn.com/problems/contains-duplicate-ii/
// Topics: 数组 哈希表

func containsNearbyDuplicate(nums []int, k int) bool {
	var flag = make(map[int]int)
	for i, n := range nums {
		if j, ok := flag[n]; ok && i-j <= k {
			return true
		}
		flag[n] = i
	}
	return false
}
