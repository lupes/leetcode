package question_441_450

// 445. 两数相加 II
// https://leetcode-cn.com/problems/add-two-numbers-ii/comments/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r1, r2 []int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			r1 = append(r1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			r2 = append(r2, l2.Val)
			l2 = l2.Next
		}
	}
	var res *ListNode
	var c, t = 0, 0
	for i1, i2 := len(r1)-1, len(r2)-1; i1 >= 0 || i2 >= 0 || c > 0; i1, i2 = i1-1, i2-1 {
		t, c = c, 0
		if i1 >= 0 {
			t += r1[i1]
		}
		if i2 >= 0 {
			t += r2[i2]
		}
		if t > 9 {
			c = 1
			t -= 10
		}
		res = &ListNode{Val: t, Next: res}
	}
	return res
}
