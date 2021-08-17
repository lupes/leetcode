package question_91_100

import (
	"sort"

	. "github.com/lupes/leetcode/common"
)

// 99. 恢复二叉搜索树
// https://leetcode-cn.com/problems/recover-binary-search-tree/
// Topics: 树 深度优先搜索

func recoverTree(root *TreeNode) {
	nodes := innerOrder(root)
	var nums = make([]int, 0, len(nodes))
	for _, node := range nodes {
		nums = append(nums, node.Val)
	}
	sort.Ints(nums)

	var a, b *TreeNode
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Val != nums[i] {
			if a == nil {
				a = nodes[i]
			} else {
				b = nodes[i]
				break
			}
		}
	}
	a.Val, b.Val = b.Val, a.Val
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
