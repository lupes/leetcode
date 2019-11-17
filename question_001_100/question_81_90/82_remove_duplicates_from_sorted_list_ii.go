package question_81_90

// 82. 删除排序链表中的重复元素 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
// Topics: 链表

func deleteDuplicates2(head *ListNode) *ListNode {
	var (
		res  = &ListNode{Next: head}
		last = res
		now  = head
	)

	for now != nil {
		for now.Next != nil && now.Val == now.Next.Val {
			now = now.Next
		}
		if last.Next == now {
			last = last.Next
		} else {
			last.Next = now.Next
		}
		now = now.Next
	}
	return res.Next
}
