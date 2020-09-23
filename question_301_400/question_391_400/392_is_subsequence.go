package question_391_400

// 392. 判断子序列
// https://leetcode-cn.com/problems/is-subsequence
// Topics: 贪心算法 二分查找 动态规划

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
