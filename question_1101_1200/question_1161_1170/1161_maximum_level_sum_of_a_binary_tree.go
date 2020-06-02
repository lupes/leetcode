package question_1161_1170

// 1161. 最大层内元素和
// https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree
// Topics: 图

import (
	. "github.com/lupes/leetcode/common"
)

func maxLevelSum(root *TreeNode) int {
	var index, max, now = 1, root.Val, []*TreeNode{root}
	for i := 1; len(now) > 0; i++ {
		var tmp = 0
		var next []*TreeNode
		for _, node := range now {
			tmp += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if tmp > max {
			index, max = i, tmp
		}
		now = next
	}
	return index
}
