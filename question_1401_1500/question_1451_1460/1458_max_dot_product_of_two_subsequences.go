package question_1451_1460

// 1458. 两个子序列的最大点积
// https://leetcode-cn.com/problems/max-dot-product-of-two-subsequences/
// Topics: 动态规划

func maxDotProduct(nums1 []int, nums2 []int) int {
	min := -1000 * 1000

	var dp = make([][]int, len(nums2)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums1)+1)
		dp[i][0] = min
	}

	for i := range dp[0] {
		dp[0][i] = min
	}

	for i, n1 := range nums2 {
		for j, n2 := range nums1 {
			r := n1 * n2

			if r+dp[i][j] > r {
				r = r + dp[i][j]
			} else if r+dp[i][j] < dp[i][j] && r < dp[i][j] {
				r = dp[i][j]
			}

			if r < dp[i+1][j] {
				r = dp[i+1][j]
			}
			if r < dp[i][j+1] {
				r = dp[i][j+1]
			}

			dp[i+1][j+1] = r
		}
	}

	return dp[len(nums2)][len(nums1)]
}
