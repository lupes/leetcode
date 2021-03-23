package question_1551_1560

// 1551. 使数组中所有元素相等的最小操作数
// https://leetcode-cn.com/problems/minimum-operations-to-make-array-equal/
// Topics: 数学

func minOperations(n int) int {
	return (n + 1) / 2 * (n - 1 + (n+1)%2) / 2
}
