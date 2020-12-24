package question_891_900

// 897. 递增顺序查找树
// https://leetcode-cn.com/problems/increasing-order-search-tree
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func increasingBST(root *TreeNode) *TreeNode {
	nums := increasingBSTDfs(root)
	return increasingBSTCreate(nums)
}

func increasingBSTDfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := increasingBSTDfs(root.Left)
	nums = append(nums, root.Val)
	nums = append(nums, increasingBSTDfs(root.Right)...)
	return nums
}

func increasingBSTCreate(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{Val: nums[0], Right: increasingBSTCreate(nums[1:])}
}
