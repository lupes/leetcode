package question_661_670

// 669. 修剪二叉搜索树
// https://leetcode-cn.com/problems/trim-a-binary-search-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	} else if root.Val < L {
		return trimBST(root.Right, L, R)
	} else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}
