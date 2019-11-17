package question_161_170

// 167. 两数之和 II - 输入有序数组
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
// Topics: 数组 双指针 二分查找

func twoSum(numbers []int, target int) []int {
	l, r, t := 0, len(numbers)-1, target
	if numbers[0] < 0 {
		t = target - numbers[0]
	}
	for r > l {
		c := (l + r) / 2
		if numbers[c] == t {
			r = c
			break
		} else if numbers[c] > t {
			r = c
		} else {
			l = c + 1
		}
	}
	l = 0
	for r > l {
		if numbers[l]+numbers[r] == target {
			break
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}
	}
	return []int{l + 1, r + 1}
}
