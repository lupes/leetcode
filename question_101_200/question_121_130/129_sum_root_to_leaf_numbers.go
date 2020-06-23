package question_121_130

import (
	. "github.com/lupes/leetcode/common"
)

// 129. 求根到叶子节点数字之和
// https://leetcode-cn.com/problems/sum-root-to-leaf-numbers
// Topics: 树 深度优先搜索

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	} else {
		return sumNumbersHelper(root.Left, num) + sumNumbersHelper(root.Right, num)
	}
}
