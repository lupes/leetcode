package question_701_710

// 705. 设计哈希集合
// https://leetcode-cn.com/problems/design-hashset
// Topics: 设计 哈希表

type MyHashSet struct {
	Base int
	Set  []*Node
}

type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor3() MyHashSet {
	return MyHashSet{
		Base: 769,
		Set:  make([]*Node, 769),
	}
}

func (this *MyHashSet) Add(key int) {
	tmp := key / this.Base
	if this.Set[tmp] == nil {
		this.Set[tmp] = &Node{Val: key}
		return
	}
	var last *Node
	for next := this.Set[tmp]; next != nil; next = next.Next {
		if next.Val == key {
			return
		}
		last = next
	}
	last.Next = &Node{Val: key}
}

func (this *MyHashSet) Remove(key int) {
	tmp := key / this.Base
	next := this.Set[tmp]
	if next == nil {
		return
	}
	if next.Val == key {
		this.Set[tmp] = next.Next
		return
	}
	var last *Node
	for ; next != nil; next = next.Next {
		if next.Val == key {
			last.Next = next.Next
			return
		}
		last = next
	}
	return
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	for next := this.Set[key/this.Base]; next != nil; next = next.Next {
		if next.Val == key {
			return true
		}
	}
	return false
}
