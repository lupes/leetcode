package offer

// 剑指 Offer II 069. 山峰数组的顶部
// https://leetcode-cn.com/problems/B1IidL/
// Topics: 数组 二分查找

func peakIndexInMountainArray(arr []int) int {
	left, right, l := 0, len(arr)-1, len(arr)
	if arr[0] > arr[1] {
		return 0
	} else if arr[l-1] > arr[l-2] {
		return l - 1
	}

	for right > left {
		c := (right + left) / 2
		if arr[c] > arr[c-1] && arr[c] > arr[c+1] {
			return c
		} else if arr[c] < arr[c+1] {
			left = c + 1
		} else {
			right = c
		}
	}
	return left
}
