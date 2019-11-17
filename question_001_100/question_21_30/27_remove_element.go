package question_0011_0020

// 27. 移除元素
// https://leetcode-cn.com/problems/remove-element
// Topics: 数组 双指针

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v == val {
			continue
		}
		nums[index] = v
		index++
	}
	return index
}
