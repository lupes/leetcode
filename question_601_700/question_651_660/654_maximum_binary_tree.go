package question_651_660

import (
	. "github.com/lupes/leetcode/common"
)

// 654. 最大二叉树
// https://leetcode-cn.com/problems/maximum-binary-tree
// Topics: 树

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, max := 0, nums[0]
	for i, n := range nums {
		if n > max {
			max = n
			index = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}
