package question_101_110

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// Topics: 树 深度优先搜索 数组

import (
	. "github.com/lupes/leetcode/common"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	var node = &TreeNode{
		Val: preorder[0],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == node.Val {
			break
		}
	}
	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return node
}
