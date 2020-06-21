package question_101_110

// 107. 二叉树的层次遍历 II
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/submissions/
// Topics: 树 广度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var next = []*TreeNode{root}
	for len(next) > 0 {
		var nextNext []*TreeNode
		var t []int
		for _, node := range next {
			t = append(t, node.Val)
			if node.Left != nil {
				nextNext = append(nextNext, node.Left)
			}
			if node.Right != nil {
				nextNext = append(nextNext, node.Right)
			}
		}
		next = nextNext
		res = append([][]int{t}, res...)
	}

	return res
}
