package question_1031_1040

// 1038. 从二叉搜索树到更大和树
// https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree
// Topics: 二叉搜索树

import (
	. "github.com/lupes/leetcode/common"
)

func bstToGst(root *TreeNode) *TreeNode {
	bstToGstHelper(root, 0)
	return root
}

func bstToGstHelper(root *TreeNode, n int) (t int) {
	if root == nil {
		return n
	}
	t = bstToGstHelper(root.Right, n) + root.Val
	root.Val = t
	t = bstToGstHelper(root.Left, t)
	return t
}
