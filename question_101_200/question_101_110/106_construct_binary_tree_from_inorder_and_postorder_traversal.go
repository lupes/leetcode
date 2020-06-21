package question_101_110

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
// Topics: 树 深度优先搜索 数组

import (
	. "github.com/lupes/leetcode/common"
)

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	var node = &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == node.Val {
			break
		}
	}
	node.Left = buildTree106(inorder[:i], postorder[:i])
	node.Right = buildTree106(inorder[i+1:], postorder[i:len(postorder)-1])
	return node
}
