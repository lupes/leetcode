package question_961_970

// 965. 单值二叉树
// https://leetcode-cn.com/problems/univalued-binary-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func isUnivalTree(root *TreeNode) bool {
	return isUnivalTreeHelper(root, root.Val)
}

func isUnivalTreeHelper(root *TreeNode, num int) bool {
	if root == nil {
		return true
	} else if root.Val != num {
		return false
	} else {
		return isUnivalTreeHelper(root.Left, num) && isUnivalTreeHelper(root.Right, num)
	}
}
