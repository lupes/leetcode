package question_01_10

import . "github.com/lupes/leetcode/common"

// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/
// Topics: 链表 数学

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	for next := res; l1 != nil || l2 != nil; next = next.Next {
		v1, v2 := 0, 0
		if l1 != nil {
			v1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			v2, l2 = l2.Val, l2.Next
		}
		if l1 != nil || l2 != nil || (v1+v2+next.Val)/10 != 0 {
			next.Next = &ListNode{Val: (v1 + v2 + next.Val) / 10}
		}
		next.Val = (v1 + v2 + next.Val) % 10
	}
	return res
}
