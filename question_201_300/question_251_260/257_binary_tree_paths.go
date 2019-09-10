package question_251_260

import "strconv"

// 257. 二叉树的所有路径
// https://leetcode-cn.com/problems/binary-tree-paths/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	s := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{s}
	}
	if root.Left != nil {
		tmp := binaryTreePaths(root.Left)
		for _, t := range tmp {
			res = append(res, s+"->"+t)
		}
	}
	if root.Right != nil {
		tmp := binaryTreePaths(root.Right)
		for _, t := range tmp {
			res = append(res, s+"->"+t)
		}
	}
	return res
}
