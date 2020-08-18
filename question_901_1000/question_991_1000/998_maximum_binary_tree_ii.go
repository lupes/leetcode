package question_991_1000

// 998. 最大二叉树 II
// https://leetcode-cn.com/problems/maximum-binary-tree-ii
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	arr := treeToArr(root)
	return createMaxTree(append(arr, val))
}

func treeToArr(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := treeToArr(root.Left)
	right := treeToArr(root.Right)
	left = append(left, root.Val)
	return append(left, right...)
}

func createMaxTree(nums []int) *TreeNode {
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
		Left:  createMaxTree(nums[:index]),
		Right: createMaxTree(nums[index+1:]),
	}
}
