package question_1301_1310

// 1302. 层数最深叶子节点的和
// https://leetcode-cn.com/problems/deepest-leaves-sum/
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func deepestLeavesSum(root *TreeNode) int {
	next, res := make([]*TreeNode, 0), 0
	for now := []*TreeNode{root}; len(now) > 0; now = next {
		res, next = 0, nil
		for _, n := range now {
			res += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
	}
	return res
}
