package question_141_150

// 147. 对链表进行插入排序
// https://leetcode-cn.com/problems/insertion-sort-list/
// Topics: 排序 链表

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res, last := &ListNode{Next: head}, head.Next
	head.Next = nil
	for last != nil {
		next := res
		for ; next.Next != nil; next = next.Next {
			if last.Val < next.Next.Val {
				break
			}
		}
		tmp := last.Next
		last.Next, next.Next, last = next.Next, last, tmp
	}
	return res.Next
}
