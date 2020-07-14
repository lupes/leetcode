package question_421_430

// 430. 扁平化多级双向链表
// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list
// Topics: 深度优先搜索 链表

type Node3 struct {
	Val   int
	Prev  *Node3
	Next  *Node3
	Child *Node3
}

func flatten(root *Node3) *Node3 {
	return nil
}
