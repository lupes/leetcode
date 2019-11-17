package question_601_610

import "fmt"

// 606. 根据二叉树创建字符串
// https://leetcode-cn.com/problems/construct-string-from-binary-tree/
// Topics: 树 字符串

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	left := tree2str(t.Left)
	right := tree2str(t.Right)
	if left == "" && right == "" {
		return fmt.Sprintf("%d", t.Val)
	} else if left != "" && right == "" {
		return fmt.Sprintf("%d(%s)", t.Val, left)
	} else {
		return fmt.Sprintf("%d(%s)(%s)", t.Val, left, right)
	}
}
