package question_531_540

// 538. 把二叉搜索树转换为累加树
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBSTHelper(root, 0)
	return root
}

func convertBSTHelper(root *TreeNode, n int) int {
	if root.Right == nil && root.Left == nil {
		root.Val += n
		return root.Val
	}
	if root.Right != nil {
		n = convertBSTHelper(root.Right, n)
	}
	root.Val += n
	if root.Left != nil {
		return convertBSTHelper(root.Left, root.Val)
	}
	return root.Val
}
