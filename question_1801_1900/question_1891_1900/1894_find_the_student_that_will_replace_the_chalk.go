package question_1891_1900

// 1894. 找到需要补充粉笔的学生编号
// https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk/
// Topics: 数组 二分查找 前缀和 模拟

func chalkReplacer(chalk []int, k int) int {
	var l = len(chalk)
	for i := 1; i < l; i++ {
		chalk[i] = chalk[i-1] + chalk[i]
	}
	k = k % (chalk[l-1])

	left, right := 0, l
	for right > left {
		center := (right + left) / 2
		if chalk[center] == k {
			return (center + 1) % l
		} else if chalk[center] < k {
			left = center + 1
		} else {
			right = center
		}
	}
	return left
}
