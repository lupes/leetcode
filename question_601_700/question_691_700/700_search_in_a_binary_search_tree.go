package question_691_700

// 700. 二叉搜索树中的搜索
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
