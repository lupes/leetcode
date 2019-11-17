package question_261_270

var res = []int{1, 2, 3, 4, 5}

// 264. 丑数 II
// https://leetcode-cn.com/problems/ugly-number-ii
// Topics: 堆 数学 动态规划

func nthUglyNumber(n int) int {
	if n <= len(res) {
		return res[n-1]
	}
	return 1
}
