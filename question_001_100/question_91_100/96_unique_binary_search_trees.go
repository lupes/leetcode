package question_91_100

// 96. 不同的二叉搜索树
// https://leetcode-cn.com/problems/unique-binary-search-trees/
// Topics: 树 动态规划

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	var nums = make([]int, n+1)
	nums[0] = 1
	for i := 1; i <= n; i++ {
		tmp := 0
		for j := 1; j <= i; j++ {
			tmp += nums[j-1] * nums[i-j]
		}
		nums[i] = tmp
	}
	return nums[n]
}
