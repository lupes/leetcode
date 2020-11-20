package question_541_550

// 543. 二叉树的直径
// https://leetcode-cn.com/problems/diameter-of-binary-tree
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func diameterOfBinaryTree(root *TreeNode) int {
	max, _ := diameterOfBinaryTreeHelper(root)
	return max
}

func diameterOfBinaryTreeHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lm, ll := diameterOfBinaryTreeHelper(root.Left)
	rm, rl := diameterOfBinaryTreeHelper(root.Right)
	max, len := lm, ll
	if rl > len {
		len = rl
	}
	if rm > max {
		max = rm
	}
	if ll+rl > max {
		max = ll + rl
	}
	return max, len + 1
}
