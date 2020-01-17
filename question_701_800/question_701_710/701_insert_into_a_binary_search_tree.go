package question_701_710

// 701. 二叉搜索树中的插入操作
// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
