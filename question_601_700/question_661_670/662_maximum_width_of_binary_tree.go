package question_661_670

// 662. 二叉树最大宽度
// https://leetcode-cn.com/problems/maximum-width-of-binary-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var now = []*TreeNode{root}
	var max = 1
	for len(now) > 0 {
		var next []*TreeNode
		for _, n := range now {
			if n == nil {
				next = append(next, nil, nil)
			} else {
				next = append(next, n.Left, n.Right)
			}
		}

		left, right := 0, len(next)
		for ; left < right && next[left] == nil; left++ {
		}
		for ; left < right && next[right-1] == nil; right-- {
		}
		now = next[left:right]
		if len(now) > max {
			max = len(now)
		}
	}
	return max
}
