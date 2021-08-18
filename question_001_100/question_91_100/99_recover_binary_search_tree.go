package question_91_100

import (
	. "github.com/lupes/leetcode/common"
)

// 99. 恢复二叉搜索树
// https://leetcode-cn.com/problems/recover-binary-search-tree/
// Topics: 树 深度优先搜索

func recoverTree(root *TreeNode) {
	nodes := innerOrder(root)

	var big, small *TreeNode
	for i := 0; i < len(nodes)-1; i++ {
		if (i == 0 && nodes[i].Val > nodes[i+1].Val) || (i > 0 && nodes[i].Val > nodes[i-1].Val && nodes[i].Val > nodes[i+1].Val) {
			big = nodes[i]
			break
		}
	}

	for i := len(nodes) - 1; i > 0; i-- {
		if (i == len(nodes)-1 && nodes[i].Val < nodes[i-1].Val) || (i < len(nodes)-1 && nodes[i].Val < nodes[i-1].Val && nodes[i].Val < nodes[i+1].Val) {
			small = nodes[i]
			break
		}
	}
	big.Val, small.Val = small.Val, big.Val
}

func innerOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := innerOrder(root.Left)
	res = append(res, root)
	res = append(res, innerOrder(root.Right)...)
	return res
}
