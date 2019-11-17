package question_141_150

import "sort"

// 148. 排序链表
// https://leetcode-cn.com/problems/sort-list/
// Topics: 排序 链表

func sortList(head *ListNode) *ListNode {
	var arr []*ListNode
	for n := head; n != nil; n = n.Next {
		arr = append(arr, n)
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].Val > arr[j].Val {
			return false
		}
		return true
	})
	var res = &ListNode{}
	next := res
	for _, n := range arr {
		next.Next = n
		next = n
		next.Next = nil
	}
	return res.Next
}
