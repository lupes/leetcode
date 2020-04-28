package question_1001_1010

// 1008. 先序遍历构造二叉树
// https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	var i = 1
	for i = 1; i < len(preorder); i++ {
		if preorder[i] > node.Val {
			break
		}
	}
	node.Left = bstFromPreorder(preorder[1:i])
	node.Right = bstFromPreorder(preorder[i:])
	return node
}
