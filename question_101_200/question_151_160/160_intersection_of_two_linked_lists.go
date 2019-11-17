package question_151_160

// 160. 相交链表
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// Topics: 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var data = make(map[*ListNode]struct{})
	for next := headA; next != nil; next = next.Next {
		data[next] = struct{}{}
	}
	for next := headB; next != nil; next = next.Next {
		if _, ok := data[next]; ok {
			return next
		}
	}
	return nil
}
