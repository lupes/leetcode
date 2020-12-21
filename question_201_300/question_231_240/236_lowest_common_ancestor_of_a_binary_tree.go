package question_231_240

// 236. 二叉树的最近公共祖先
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
// Topics: 树
import . "github.com/lupes/leetcode/common"

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	_, _, res := lowestCommonAncestorHelper(root, p, q)
	return res
}

func lowestCommonAncestorHelper(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	ep, eq := root.Val == p.Val, root.Val == q.Val

	lep, leq, ln := lowestCommonAncestorHelper(root.Left, p, q)
	if lep && leq {
		return true, true, ln
	}
	if (leq && ep) || (lep && eq) {
		return true, true, root
	}
	rep, req, rn := lowestCommonAncestorHelper(root.Right, p, q)
	if rep && req {
		return true, true, rn
	}
	if (req && ep) || (rep && eq) || (lep && req) || (leq && rep) {
		return true, true, root
	}
	return ep || lep || rep, eq || leq || req, nil
}
