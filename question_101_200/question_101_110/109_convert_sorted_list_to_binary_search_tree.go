package question_101_110

// 109. 有序链表转换二叉搜索树
// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
// Topics: 深度优先搜索 链表

import (
	. "github.com/lupes/leetcode/common"
)

func sortedListToBST(head *ListNode) *TreeNode {
	var values []int
	for ; head != nil; head = head.Next {
		values = append(values, head.Val)
	}
	return sortedListToBSTHelper(values)
}

func sortedListToBSTHelper(values []int) *TreeNode {
	l := len(values)
	if l == 0 {
		return nil
	} else if l == 1 {
		return &TreeNode{Val: values[0]}
	} else {
		return &TreeNode{
			Val:   values[l/2],
			Left:  sortedListToBSTHelper(values[:l/2]),
			Right: sortedListToBSTHelper(values[l/2+1:]),
		}
	}
}
