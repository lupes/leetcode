package question_1451_1460

// 1457. 二叉树中的伪回文路径
// https://leetcode-cn.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
// Topics: 位 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func pseudoPalindromicPaths(root *TreeNode) int {
	return pseudoPalindromicPathsHelper(root, 0)
}

func pseudoPalindromicPathsHelper(root *TreeNode, t int) int {
	sum := 0
	t ^= 1 << root.Val
	if root.Left == nil && root.Right == nil {
		switch t {
		case 0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512:
			return 1
		default:
			return 0
		}
	}
	if root.Left != nil {
		sum += pseudoPalindromicPathsHelper(root.Left, t)
	}
	if root.Right != nil {
		sum += pseudoPalindromicPathsHelper(root.Right, t)
	}
	return sum
}
