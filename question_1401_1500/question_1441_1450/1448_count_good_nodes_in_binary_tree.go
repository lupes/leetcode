package question_1441_1450

// 1448. 统计二叉树中好节点的数目
// https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree/
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func goodNodes(root *TreeNode) int {
	var max = root.Val
	return goodNodesHelper(root, max)
}

func goodNodesHelper(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	var res = 0
	if root.Val >= max {
		res++
		max = root.Val
	}
	res += goodNodesHelper(root.Left, max)
	res += goodNodesHelper(root.Right, max)
	return res
}
