package question_1601_1610

// 1609. 奇偶树
// https://leetcode-cn.com/problems/even-odd-tree/
// Topics: 树

import (
	. "github.com/lupes/leetcode/common"
)

func isEvenOddTree(root *TreeNode) bool {
	var now, level = []*TreeNode{root}, 0
	for len(now) > 0 {
		var next []*TreeNode
		var tmp = now[0].Val
		if level&1 == 0 {
			tmp -= 1
		} else {
			tmp += 1
		}

		for _, node := range now {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

			if level&1 == 0 { // 偶数行
				if node.Val&1 == 0 || node.Val <= tmp {
					return false
				}
			} else {
				if node.Val&1 == 1 || node.Val >= tmp {
					return false
				}
			}
			tmp = node.Val
		}
		level++
		now = next
	}
	return true
}
