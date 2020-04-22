package question_991_1000

// 993. 二叉树的堂兄弟节点
// https://leetcode-cn.com/problems/cousins-in-binary-tree
// Topics: 树 广度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func isCousins(root *TreeNode, x int, y int) bool {
	var xh, xp, yh, yp int
	if root.Val == x {
		xh, xp = 0, 0
	} else {
		xh, xp = isCousinsHelper(root, x, 1)
	}

	if root.Val == y {
		yh, yp = 0, 0
	} else {
		yh, yp = isCousinsHelper(root, y, 1)
	}

	return xh == yh && xp != yp
}

func isCousinsHelper(root *TreeNode, x, h int) (xh int, xp int) {
	if root.Left != nil {
		if root.Left.Val == x {
			return h, root.Val
		}
		xh, xp = isCousinsHelper(root.Left, x, h+1)
		if xh != -1 {
			return xh, xp
		}
	}
	if root.Right != nil {
		if root.Right.Val == x {
			return h, root.Val
		}
		xh, xp = isCousinsHelper(root.Right, x, h+1)
		if xh != -1 {
			return xh, xp
		}
	}
	return -1, 0
}
