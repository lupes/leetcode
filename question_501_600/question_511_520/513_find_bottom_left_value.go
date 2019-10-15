package main

// 513. 找树左下角的值
// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var current, res = []*TreeNode{root}, root.Val
	for len(current) > 0 {
		var next []*TreeNode
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res, current = current[0].Val, next
	}
	return res
}
