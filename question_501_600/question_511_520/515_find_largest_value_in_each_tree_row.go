package qustion_511_520

// 515. 在每个树行中找最大值
// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
// Topics: 树 深度优先搜索 广度优先搜索

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var current = []*TreeNode{root}
	for len(current) > 0 {
		var next []*TreeNode
		var max = current[0].Val
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			if node.Val > max {
				max = node.Val
			}
		}
		res = append(res, max)
		current = next
	}
	return res
}
