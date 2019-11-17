package question_221_230

import (
	"math"
)

// 222. 完全二叉树的节点个数
// https://leetcode-cn.com/problems/count-complete-tree-nodes/
// Topics: 树 二分查找

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	level, res := 0, 0
	for n := root; n != nil; n = n.Left {
		if n.Left != nil {
			res += int(math.Pow(2, float64(level)))
		}
		level++
	}
	if level == 1 {
		return 1
	}
	sum := int(math.Pow(2, float64(level-1)))
	left, right := 0, sum
	for j := level; j > 0; j-- {
		center := (left + right) / 2
		left2, right2 := left, right
		next := root
		for i := 1; i < j; i++ {
			if center > (left2+right2)/2 {
				next = next.Right
				left2 = (left2 + right2) / 2
			} else {
				next = next.Left
				right2 = (left2 + right2) / 2
			}
		}
		if left == right {
			if next == nil {
				res += left - 1
			} else {
				res += left
			}
			break
		}
		if next != nil {
			left = (left+right)/2 + 1
			root = root.Right
		} else {
			right = (left + right) / 2
			root = root.Left
		}
	}

	return res
}
