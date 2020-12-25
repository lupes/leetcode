package question_891_900

// 897. 递增顺序查找树
// https://leetcode-cn.com/problems/increasing-order-search-tree
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func increasingBST(root *TreeNode) *TreeNode {
	parent := &TreeNode{}
	increasingBSTDfs(root, parent)
	return parent.Right
}

func increasingBSTDfs(root, parent *TreeNode) *TreeNode {
	if root == nil {
		return parent
	}
	next := increasingBSTDfs(root.Left, parent)
	next.Right = &TreeNode{Val: root.Val}
	return increasingBSTDfs(root.Right, next.Right)
}
