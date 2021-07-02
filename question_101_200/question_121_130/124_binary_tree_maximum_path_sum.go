package question_121_130

import (
	. "github.com/lupes/leetcode/common"
)

// 124. 二叉树中的最大路径和
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
// Topics: 树 深度优先搜索

func maxPathSum(root *TreeNode) int {
	_, max := maxPathSumHelp(root)
	return max
}

func maxFunc(max int, args ...int) int {
	for _, arg := range args {
		if arg > max {
			max = arg
		}
	}
	return max
}

func maxPathSumHelp(root *TreeNode) (int, int) {
	if root == nil {
		return -1001, -1001
	}

	leftMax, leftSum := maxPathSumHelp(root.Left)
	rightMax, rightSum := maxPathSumHelp(root.Right)

	var max = maxFunc(root.Val, leftMax+root.Val, rightMax+root.Val)
	var sum = maxFunc(root.Val, leftSum, rightSum, leftMax+root.Val, rightMax+root.Val, leftMax+root.Val+rightMax)
	return max, sum
}
