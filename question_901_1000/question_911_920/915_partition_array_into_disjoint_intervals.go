package question_911_920

// 915. 分割数组
// https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals
// Topics: 数组

func partitionDisjoint(A []int) int {
	max, min, len := 0, 10000000, len(A)
	var flag = make([][2]int, len)
	for i := 0; i < len; i++ {
		if A[i] > max {
			max = A[i]
		}
		flag[i][0] = max

		if min > A[len-1-i] {
			min = A[len-1-i]
		}
		flag[len-1-i][1] = min
	}
	for i := 0; i < len-1; i++ {
		if flag[i][0] <= flag[i+1][1] {
			return i + 1
		}
	}
	return 0
}
