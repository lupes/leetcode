package question_1321_1330

import (
	. "github.com/lupes/leetcode/common"
)

// 1325. 删除给定值的叶子节点
// https://leetcode-cn.com/problems/delete-leaves-with-a-given-value/
// Topics: 字符串

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
