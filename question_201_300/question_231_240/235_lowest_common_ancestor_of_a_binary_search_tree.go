package question_231_240

// 235. 二叉搜索树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
// Topics: 树

import . "github.com/lupes/leetcode/common"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
