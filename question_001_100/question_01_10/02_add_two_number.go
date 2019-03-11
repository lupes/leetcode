package question_01_10

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) Test() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	next := res
	height := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		val := (v1 + v2 + height) % 10
		height = (v1 + v2 + height) / 10
		next.Next = &ListNode{Val: val}
		next = next.Next
	}
	if height != 0 {
		next.Next = &ListNode{Val: 1}
	}
	return res.Next
}
