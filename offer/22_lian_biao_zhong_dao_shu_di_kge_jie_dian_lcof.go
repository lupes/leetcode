package offer

import (
	. "github.com/lupes/leetcode/common"
)

// 剑指 Offer 22. 链表中倒数第k个节点
// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
// Topics: 链表 双指针

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast = head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		head = head.Next
	}
	return head
}
