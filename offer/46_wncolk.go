package offer

import (
	. "github.com/lupes/leetcode/common"
)

// 剑指 Offer II 046. 二叉树的右侧视图
// https://leetcode-cn.com/problems/WNC0Lk/
// Topics: 树 深度优先搜索 广度优先搜索 二叉树

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var next = []*TreeNode{root}
	for len(next) > 0 {
		res = append(res, next[len(next)-1].Val)
		var tmp []*TreeNode
		for _, n := range next {
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		next = tmp
	}
	return res
}
