package question_651_660

// 658. 找到 K 个最接近的元素
// https://leetcode-cn.com/problems/find-k-closest-elements
// Topics: 二分查找

func findClosestElements(arr []int, k int, x int) []int {
	var l, left, right, ct = len(arr), 0, len(arr), 0
	if l == k {
		return arr
	}
	for right > left+1 {
		ct = (left + right) / 2
		if arr[ct] > x {
			right = ct - 1
		} else {
			left = ct
		}
	}
	left = right - k - 1
	if left < 0 {
		left = 0
	}
	right = right + k + 1
	if right > l {
		right = l
	}
	for right > left+k {
		ld := arr[left] - x
		rd := arr[right-1] - x
		if ld*ld > rd*rd {
			left++
		} else {
			right--
		}
	}
	return arr[left:right]
}
