package question_501_510

import (
	. "github.com/lupes/leetcode/common"
)

// 501. 二叉搜索树中的众数
// https://leetcode-cn.com/problems/find-mode-in-binary-search-tree
// Topics: 树

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var max, now, times = 0, 0, 0
	var res []int
	var nodes = []*TreeNode{root}
	for len(nodes) > 0 {
		top := nodes[len(nodes)-1]
		if top.Left != nil {
			nodes = append(nodes, top.Left)
			top.Left = nil
		} else {
			if top.Val == now {
				times++
			} else {
				times = 1
				now = top.Val
			}
			if times > max {
				res = []int{now}
				max = times
			} else if times == max {
				res = append(res, now)
			}
			nodes = nodes[:len(nodes)-1]
			if top.Right != nil {
				nodes = append(nodes, top.Right)
				top.Right = nil
			}
		}
	}
	return res
}
