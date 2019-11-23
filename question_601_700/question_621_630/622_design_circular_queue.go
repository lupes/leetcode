package question_621_630

// 622. 设计循环队列
// https://leetcode-cn.com/problems/design-circular-queue
// Topics: 设计 队列

type MyCircularQueue struct {
	queue []int
	head  int
	end   int
	len   int
	k     int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		head:  -1,
		end:   0,
		len:   0,
		k:     k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head++
	if this.head == this.k {
		this.head = 0
	}
	this.queue[this.head] = value
	this.len++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.end++
	if this.end == this.k {
		this.end = 0
	}
	this.len--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.end]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.k
}
