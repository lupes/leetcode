package question_951_960

// 951. 翻转等价二叉树
// https://leetcode-cn.com/problems/flip-equivalent-binary-trees
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
