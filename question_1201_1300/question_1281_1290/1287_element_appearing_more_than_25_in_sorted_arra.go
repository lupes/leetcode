package question_1281_1290

// 1287. 有序数组中出现次数超过25%的元素
// https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
// Topics: 数组

func findSpecialInteger(arr []int) int {
	for i, l := 0, len(arr)/4; i < len(arr); i++ {
		if arr[i] == arr[i+l] {
			return arr[i]
		}
	}
	return 0
}
