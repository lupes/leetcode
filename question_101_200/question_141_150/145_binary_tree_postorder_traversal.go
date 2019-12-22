package question_141_150

import (
	"fmt"

	. "github.com/lupes/leetcode/common"
)

// 145. 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal
// Topics: 栈 树

func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stacks := []*TreeNode{root}
	for len(stacks) > 0 {
		size := len(stacks) - 1
		node := stacks[size]
		stacks = stacks[:size]
		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}

		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}

		//res = append(res, node.Val)
		res = append([]int{node.Val}, res...)
	}
	fmt.Printf("%v\n", res)
	return res
}
