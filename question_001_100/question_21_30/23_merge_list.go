package question_0011_0020

type ListNode struct {
	Val  int
	Next *ListNode
}

// 23. 合并K个排序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	for _, node := range lists {
		if node == nil {
			continue
		}
		if head == nil {
			head = node
			continue
		}
		for {
			old := head
			if head.Val > node.Val {
				head = node
				node = node.Next
				head.Next = old
			} else {
				for {
					if old.Val <= node.Val {
						if old.Next != nil {
							if old.Next.Val >= node.Val {
								tmp := old.Next
								old.Next = node
								node = node.Next
								old.Next.Next = tmp
								break
							} else {
								old = old.Next
							}
						} else {
							old.Next = node
							node = node.Next
							old.Next.Next = nil
							break
						}
					}
				}
			}

			if node == nil {
				break
			}
		}
	}
	return head
}
