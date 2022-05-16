package mian_shi

import (
	. "github.com/lupes/leetcode/common"
)

// 面试题 04.06. 后继者
// https://leetcode.cn/problems/successor-lcci/
// Topics: 树 深度优先搜索 二叉搜索树 二叉树

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	n, _ := inorderSuccessorHelper(root, p, false)
	return n
}

func inorderSuccessorHelper(root *TreeNode, p *TreeNode, find bool) (*TreeNode, bool) {
	if root == nil {
		return nil, find
	}

	if root.Val == p.Val {
		return inorderSuccessorHelper(root.Right, p, true)
	} else if root.Val < p.Val {
		return inorderSuccessorHelper(root.Right, p, false)
	} else {
		node, find := inorderSuccessorHelper(root.Left, p, find)
		if find && node != nil {
			return node, find
		} else if find {
			return root, find
		}
	}
	return nil, false
}
