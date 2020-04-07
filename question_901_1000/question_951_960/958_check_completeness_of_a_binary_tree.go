package question_951_960

// 958. 二叉树的完全性检验
// https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree
// Topics: 树

import (
	"fmt"

	. "github.com/lupes/leetcode/common"
)

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var now = []*TreeNode{root}

	for col := uint(0); ; col++ {
		num, flag := 0, false
		var next []*TreeNode
		for _, node := range now {
			if node.Left != nil && !flag {
				next = append(next, node.Left)
				num++
			} else if node.Left != nil && flag {
				return false
			} else {
				flag = true
			}
			if node.Right != nil && !flag {
				next = append(next, node.Right)
				num++
			} else if node.Right != nil && flag {
				return false
			} else {
				flag = true
			}
		}
		if len(now) <= 1<<col && num == 0 {
			return true
		} else if len(now) < 1<<col && num != 0 {
			return false
		} else {
			now = next
		}
		fmt.Printf("%+v\n", now)
	}
}
