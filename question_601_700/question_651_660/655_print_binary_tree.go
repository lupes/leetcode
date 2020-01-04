package question_651_660

import (
	"fmt"

	. "github.com/lupes/leetcode/common"
)

// 655. 输出二叉树
// https://leetcode-cn.com/problems/print-binary-tree
// Topics: 树

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	height := depth(root)
	weight := 1<<height - 1
	var res = make([][]string, height)
	for i := 0; i < height; i++ {
		res[i] = make([]string, weight)
	}
	printTreeHelper(root, res, height, 0, 1<<(height-1)-1)
	return res
}

func printTreeHelper(root *TreeNode, res [][]string, height, i, j int) {
	res[i][j] = fmt.Sprintf("%d", root.Val)
	if root.Left != nil {
		printTreeHelper(root.Left, res, height, i+1, j-(1<<(height-i-2)))
	}
	if root.Right != nil {
		printTreeHelper(root.Right, res, height, i+1, j+(1<<(height-i-2)))
	}
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
