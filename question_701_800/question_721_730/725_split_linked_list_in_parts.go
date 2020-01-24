package question_721_730

// 725. 分隔链表
// https://leetcode-cn.com/problems/split-linked-list-in-parts
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func splitListToParts(root *ListNode, k int) []*ListNode {
	if k == 1 {
		return []*ListNode{root}
	}
	l := 0
	for next := root; next != nil; next = next.Next {
		l++
	}
	var res []*ListNode
	a, b, next := l/k, l%k, root
	for i := 0; i < k; i++ {
		if next == nil {
			res = append(res, nil)
		} else {
			res = append(res, next)
			c := a
			if i < b {
				c++
			}
			for j := 1; j < c && next != nil; j++ {
				next = next.Next
			}
			if next != nil {
				next, next.Next = next.Next, nil
			}
		}
	}
	return res
}
