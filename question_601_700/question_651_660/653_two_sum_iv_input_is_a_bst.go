package question_651_660

import (
	. "github.com/lupes/leetcode/common"
)

// 653. 两数之和 IV - 输入 BST
// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
// Topics: 树

func findTarget(root *TreeNode, k int) bool {
	var flags = make(map[int]struct{})
	return findTargetHelper(root, flags, k)
}

func findTargetHelper(root *TreeNode, flags map[int]struct{}, k int) bool {
	if root == nil {
		return false
	}
	if _, ok := flags[k-root.Val]; ok {
		return true
	}
	flags[root.Val] = struct{}{}
	return findTargetHelper(root.Left, flags, k) || findTargetHelper(root.Right, flags, k)
}
