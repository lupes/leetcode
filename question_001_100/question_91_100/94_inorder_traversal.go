package question_91_100

import . "github.com/lupes/leetcode/common"

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	} else {
		res = append(res, inorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}
