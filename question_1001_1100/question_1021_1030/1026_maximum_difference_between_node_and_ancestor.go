package question_1021_1030

// 1026. 节点与其祖先之间的最大差值
// https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffHelper(root, root.Val, root.Val)
}

func maxAncestorDiffHelper(root *TreeNode, max, min int) int {
	if root == nil {
		return 0
	}
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	res := max - min
	tmp := maxAncestorDiffHelper(root.Left, max, min)
	if tmp > res {
		res = tmp
	}
	tmp = maxAncestorDiffHelper(root.Right, max, min)
	if tmp > res {
		res = tmp
	}
	return res
}
