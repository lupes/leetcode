package question_1661_1670

// 1663. 具有给定数值的最小字符串
// https://leetcode-cn.com/problems/smallest-string-with-a-given-numeric-value/
// Topics: 贪心算法

func getSmallestString(n int, k int) string {
	var res, sum = make([]byte, n), n * 26

	for i := 0; i < n; i++ {
		if sum-25 > k {
			res[i], sum = 'a', sum-25
		} else {
			res[i], sum = 'z'-byte(sum-k), k
		}
	}
	return string(res)
}
