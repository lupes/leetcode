package question_981_990

// 988. 从叶结点开始的最小字符串
// https://leetcode-cn.com/problems/smallest-string-starting-from-leaf
// Topics: 树 深度优先搜索

import (
	. "github.com/lupes/leetcode/common"
)

func smallestFromLeaf(root *TreeNode) string {
	return smallestFromLeafHelper(root, "", "")
}

func smallestFromLeafHelper(root *TreeNode, suffix, min string) string {
	t := string('a'+root.Val) + suffix
	if root.Left == nil && root.Right == nil {
		if (min != "" && t < min) || min == "" {
			min = t
		}
	}
	if root.Left != nil {
		min = smallestFromLeafHelper(root.Left, t, min)
	}
	if root.Right != nil {
		min = smallestFromLeafHelper(root.Right, t, min)
	}
	return min
}
