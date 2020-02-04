package question_781_790

// 783. 二叉搜索树结点最小距离
// https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
// Topics: 树 递归

import (
	. "github.com/lupes/leetcode/common"
)

func minDiffInBST(root *TreeNode) int {
	var arr []int
	minDiffInBSTHelper(root, &arr)
	var min = 1<<63 - 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
	}
	return min
}

func minDiffInBSTHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	minDiffInBSTHelper(root.Left, res)
	*res = append(*res, root.Val)
	minDiffInBSTHelper(root.Right, res)
}
