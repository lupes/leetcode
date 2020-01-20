package question_701_710

import (
	"fmt"
)

// 707. 设计链表
// https://leetcode-cn.com/problems/design-linked-list
// Topics: 设计 链表

func (this MyLinkedList) String() string {
	return fmt.Sprintf("%d %s", this.Val, this.Next)
}

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
	Prev *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor5() MyLinkedList {
	return MyLinkedList{Val: 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.findIndex(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (this *MyLinkedList) addFirst(val int) {
	now := &MyLinkedList{Val: val}
	this.Next = now
	this.Prev = now
}

func (this *MyLinkedList) findIndex(index int) *MyLinkedList {
	var next = this.Next
	for i := 0; i < index && next != nil; i++ {
		next = next.Next
	}
	return next
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Val == 0 {
		this.addFirst(val)
	} else {
		now := &MyLinkedList{Val: val}
		now.Next = this.Next
		this.Next.Prev = now
		this.Next = now
	}
	this.Val++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Val == 0 {
		this.addFirst(val)
	} else {
		now := &MyLinkedList{Val: val}
		now.Prev = this.Prev
		now.Prev.Next = now
		this.Prev = now
	}
	this.Val++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	} else if index == this.Val {
		this.AddAtTail(val)
		return
	} else if index > this.Val {
		return
	}
	node := this.findIndex(index)
	now := &MyLinkedList{Val: val}
	now.Next = node
	now.Prev = node.Prev
	node.Prev.Next = now
	node.Prev = now
	this.Val++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Val || this.Val == 0 {
		return
	}
	if this.Val == 1 {
		this.Next = nil
		this.Prev = nil
	} else if index == 0 {
		this.Next = this.Next.Next
		this.Next.Prev = nil
	} else if this.Val == index+1 {
		this.Prev = this.Prev.Prev
		this.Prev.Next = nil
	} else {
		node := this.findIndex(index)
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	this.Val--
}
