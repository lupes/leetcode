package question_1301_1310

// 1304. 和为零的N个唯一整数
// https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
// Topics: 数组

func sumZero(n int) []int {
	var res = make([]int, 0, n)
	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}
	if n%2 == 1 {
		res = append(res, 0)
	}
	return res
}
