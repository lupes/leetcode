package question_441_450

// 450. 删除二叉搜索树中的节点
// https://leetcode-cn.com/problems/delete-node-in-a-bst
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else {
			now := root.Left
			for now.Right != nil {
				now = now.Right
			}
			now.Right = root.Right
			return root.Left
		}
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
