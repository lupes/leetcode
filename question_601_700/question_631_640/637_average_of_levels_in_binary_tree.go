package question_631_640

import (
	. "github.com/lupes/leetcode/common"
)

// 637. 二叉树的层平均值
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
// Topics: 树

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var res []float64
	var now = []*TreeNode{root}
	for len(now) > 0 {
		var next []*TreeNode
		var tmp = 0
		for _, node := range now {
			tmp += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, float64(tmp)/float64(len(now)))
		now = next
	}
	return res
}
