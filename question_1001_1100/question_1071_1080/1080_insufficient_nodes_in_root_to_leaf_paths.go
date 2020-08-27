package question_1071_1080

// 1080. 根到叶路径上的不足节点
// https://leetcode-cn.com/problems/insufficient-nodes-in-root-to-leaf-paths
// Topics: 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	root, _ = sufficientSubsetHelp(root, limit)
	return root
}

func sufficientSubsetHelp(root *TreeNode, limit int) (*TreeNode, int) {
	limit -= root.Val
	if root.Left != nil && root.Right == nil {
		root.Left, limit = sufficientSubsetHelp(root.Left, limit)
	} else if root.Left == nil && root.Right != nil {
		root.Right, limit = sufficientSubsetHelp(root.Right, limit)
	} else if root.Left != nil && root.Right != nil {
		var min, t int
		root.Left, min = sufficientSubsetHelp(root.Left, limit)
		root.Right, t = sufficientSubsetHelp(root.Right, limit)
		if t < min {
			min = t
		}
		limit = min
	}
	if limit > 0 {
		return nil, limit
	}
	return root, limit
}
