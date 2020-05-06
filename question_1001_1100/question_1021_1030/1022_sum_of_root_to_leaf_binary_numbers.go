package question_1021_1030

// 1022. 从根到叶的二进制数之和
// https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func sumRootToLeaf(root *TreeNode) int {
	return sumRootToLeafDFS(root, 0)
}

func sumRootToLeafDFS(root *TreeNode, num int) int {
	var tmp = (num<<1 + root.Val) % 1000000007
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	var res = 0
	if root.Left != nil {
		res += sumRootToLeafDFS(root.Left, tmp)
	}
	if root.Right != nil {
		res += sumRootToLeafDFS(root.Right, tmp)
	}
	return res
}
