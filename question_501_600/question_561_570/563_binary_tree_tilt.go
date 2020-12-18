package question_561_570

// 563. 二叉树的坡度
// https://leetcode-cn.com/problems/binary-tree-tilt
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func findTilt(root *TreeNode) int {
	_, r := findTiltHelper(root)
	return r
}

func findTiltHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l1, l2 := findTiltHelper(root.Left)
	r1, r2 := findTiltHelper(root.Right)
	t := l1 - r1
	if t < 0 {
		t *= -1
	}
	return root.Val + l1 + r1, t + l2 + r2
}
