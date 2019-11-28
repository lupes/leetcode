package question_611_620

import (
	. "github.com/lupes/leetcode/common"
)

// 617. 合并二叉树
// https://leetcode-cn.com/problems/merge-two-binary-trees
// Topics: 树

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
