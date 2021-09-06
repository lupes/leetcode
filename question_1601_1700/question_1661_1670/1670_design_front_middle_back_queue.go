package question_1661_1670

// 1670. 设计前中后队列
// https://leetcode-cn.com/problems/design-front-middle-back-queue/
// Topics: 设计 队列 数组 链表 数据流

type Node struct {
	Val  int
	Next *Node
}

type FrontMiddleBackQueue struct {
	Len  int
	Head *Node
}

func Constructor() FrontMiddleBackQueue {
	this := FrontMiddleBackQueue{
		Len:  0,
		Head: &Node{},
	}
	return this
}

func (this *FrontMiddleBackQueue) addOne(front, next *Node, val int) {
	node := &Node{
		Next: next,
		Val:  val,
	}
	front.Next = node
	this.Len += 1
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.addOne(this.Head, this.Head.Next, val)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	next := this.Head
	for i := 0; i < this.Len/2; next, i = next.Next, i+1 {
	}
	this.addOne(next, next.Next, val)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	next := this.Head
	for ; next.Next != nil; next = next.Next {
	}
	this.addOne(next, nil, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.Len == 0 {
		return -1
	}

	tmp := this.Head.Next
	this.Head.Next = tmp.Next
	this.Len--
	return tmp.Val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.Len == 0 {
		return -1
	}
	mid := this.Len / 2
	if this.Len&1 == 0 {
		mid -= 1
	}

	next := this.Head
	for i := 0; i < mid; next, i = next.Next, i+1 {
	}
	tmp := next.Next
	next.Next = tmp.Next
	this.Len--
	return tmp.Val
}
func (this *FrontMiddleBackQueue) PopBack() int {
	if this.Len == 0 {
		return -1
	}

	next := this.Head
	for ; next.Next.Next != nil; next = next.Next {
	}
	tmp := next.Next
	next.Next = nil
	this.Len--
	return tmp.Val
}
