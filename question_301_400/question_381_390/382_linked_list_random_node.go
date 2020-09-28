package question_381_390

// 382. 链表随机节点
// https://leetcode-cn.com/problems/linked-list-random-node
// Topics: 蓄水池抽样

import (
	"math/rand"

	. "github.com/lupes/leetcode/common"
)

type Solution2 struct {
	head  *ListNode
	count int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor2(head *ListNode) Solution2 {
	count := 0
	for next := head; next != nil; next = next.Next {
		count += 1
	}
	return Solution2{
		head:  head,
		count: count,
	}
}

/** Returns a random node's value. */
func (this *Solution2) GetRandom() int {
	n := rand.Intn(this.count)
	for i, next := 0, this.head; next != nil; i, next = i+1, next.Next {
		if i == n {
			return next.Val
		}
	}
	return 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
