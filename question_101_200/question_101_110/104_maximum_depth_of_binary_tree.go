package question_101_110

// 104. 二叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	max := leftDepth
	if rightDepth > leftDepth {
		max = rightDepth
	}
	return max + 1
}
