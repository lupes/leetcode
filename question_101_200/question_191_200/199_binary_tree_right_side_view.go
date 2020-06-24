package question_191_200

// 199. 二叉树的右视图
// https://leetcode-cn.com/problems/binary-tree-right-side-view
// Topics: 树 深度优先搜索 广度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var now = []*TreeNode{root}
	var res []int
	for len(now) > 0 {
		res = append(res, now[len(now)-1].Val)
		var next []*TreeNode
		for _, node := range now {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		now = next
	}
	return res
}
