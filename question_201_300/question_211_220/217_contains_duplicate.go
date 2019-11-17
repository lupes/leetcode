package question_211_220

// 217. 存在重复元素
// https://leetcode-cn.com/problems/contains-duplicate/
// Topics: 数组 哈希表

func containsDuplicate(nums []int) bool {
	var flag = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := flag[n]; ok {
			return true
		}
		flag[n] = struct{}{}
	}
	return false
}
