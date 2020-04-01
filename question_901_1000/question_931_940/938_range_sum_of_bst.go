package question_931_940

// 938. 二叉搜索树的范围和
// https://leetcode-cn.com/problems/range-sum-of-bst
// Topics: 树 递归

import (
	. "github.com/lupes/leetcode/common"
)

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	var res = 0
	if root.Val < L {
		res += rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		res += rangeSumBST(root.Left, L, R)
	} else {
		res = root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	}
	return res
}
