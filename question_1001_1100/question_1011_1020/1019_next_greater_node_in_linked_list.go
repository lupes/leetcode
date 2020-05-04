package question_1011_1020

// 1019. 链表中的下一个更大节点
// https://leetcode-cn.com/problems/next-greater-node-in-linked-list
// Topics: 栈 链表

import (
	. "github.com/lupes/leetcode/common"
)

func nextLargerNodes(head *ListNode) []int {
	var stack, res, i = make([][2]int, 0), make([]int, 0), 0
	for head != nil {
		res = append(res, 0)
		if len(stack) > 0 {
			for len(stack) > 0 {
				l := len(stack)
				if head.Val > stack[l-1][1] {
					res[stack[l-1][0]] = head.Val
					stack = stack[:l-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, [2]int{i, head.Val})
		head, i = head.Next, i+1
	}
	return res
}
