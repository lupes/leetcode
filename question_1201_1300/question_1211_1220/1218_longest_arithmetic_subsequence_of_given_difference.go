package question_1211_1220

// 1218. 最长定差子序列
// https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference
// Topics: 数学 哈希表 动态规划

func longestSubsequence(arr []int, difference int) int {
	var m = make(map[int]int, len(arr))
	var res = 0
	for _, n := range arr {
		m[n] = m[n-difference] + 1
		if m[n] > res {
			res = m[n]
		}
	}
	return res
}
