package question_811_820

// 814. 二叉树剪枝
// https://leetcode-cn.com/problems/binary-tree-pruning
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
