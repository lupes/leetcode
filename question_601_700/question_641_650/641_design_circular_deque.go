package question_641_650

// 641. 设计循环双端队列
// https://leetcode-cn.com/problems/design-circular-deque
// Topics: 设计 队列

type MyCircularDeque struct {
	queue []int
	size  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		size: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if !this.IsFull() {
		this.queue = append([]int{value}, this.queue...)
		return true
	}
	return false
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if !this.IsFull() {
		this.queue = append(this.queue, value)
		return true
	}
	return false
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[1:]
		return true
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[:len(this.queue)-1]
		return true
	}
	return false
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.queue[0]
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.queue[len(this.queue)-1]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.queue) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.queue) == this.size
}
