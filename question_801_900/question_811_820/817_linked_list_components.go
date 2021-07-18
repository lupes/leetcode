package question_811_820

// 817. 链表组件
// https://leetcode-cn.com/problems/linked-list-components
// Topics: 链表

import (
	. "github.com/lupes/leetcode/common"
)

func numComponents(head *ListNode, G []int) int {
	var cnt, flag = 0, make(map[int]bool, len(G))
	for _, g := range G {
		flag[g] = true
	}
	for next := head; next != nil; next = next.Next {
		if flag[next.Val] && (next.Next == nil || !flag[next.Next.Val]) {
			cnt++
		}
	}
	return cnt
}
