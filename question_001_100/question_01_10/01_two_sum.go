package question_01_10

// 1. 两数之和
// https://leetcode-cn.com/problems/two-sum/
// Topics: 数组 哈希表

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, ok := m[target-num]
		if ok && i != j {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{0, 0}
}
