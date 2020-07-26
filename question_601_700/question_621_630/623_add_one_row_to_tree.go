package question_621_630

import (
	. "github.com/lupes/leetcode/common"
)

// 623. 在二叉树中增加一行
// https://leetcode-cn.com/problems/add-one-row-to-tree
// Topics: 树

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	addOneRowHelper(root, v, d-1)
	return root
}

func addOneRowHelper(root *TreeNode, v int, d int) {
	if root == nil {
		return
	}
	if d == 1 {
		root.Left = &TreeNode{Val: v, Left: root.Left}
		root.Right = &TreeNode{Val: v, Right: root.Right}
		return
	} else {
		addOneRowHelper(root.Left, v, d-1)
		addOneRowHelper(root.Right, v, d-1)
	}
}
