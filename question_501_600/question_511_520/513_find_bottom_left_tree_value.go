package qustion_511_520

// 513. 找树左下角的值
// https://leetcode-cn.com/problems/find-bottom-left-tree-value/
// Topics: 树 深度优先搜索 广度优先搜索

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
