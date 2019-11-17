package question_91_100

import . "github.com/lupes/leetcode/common"

// 100. 相同的树
// https://leetcode-cn.com/problems/same-tree/
// Topics: 树 深度优先搜索

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
