package question_231_240

import (
	"fmt"
)

// 232. 用栈实现队列
// https://leetcode-cn.com/problems/implement-queue-using-stacks
// Topics: 栈 设计

type stack []int

func (self *stack) push(v int) {
	*self = append(*self, v)
}

func (self *stack) pop() int {
	if !self.empty() {
		v := (*self)[0]
		*self = (*self)[1:]
		return v
	}
	return -1
}

func (self *stack) peek() int {
	if !self.empty() {
		return (*self)[0]
	}
	return -1
}

func (self stack) size() int {
	return len(self)
}

func (self stack) empty() bool {
	return len(self) == 0
}

type MyQueue struct {
	left  stack
	right stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if this.left.empty() {
		for !this.right.empty() {
			this.left.push(this.right.pop())
		}
	}
	this.left.push(x)
}

func (this *MyQueue) Pop() int {
	if this.right.empty() {
		for !this.left.empty() {
			this.right.push(this.left.pop())
		}
	}
	return this.right.pop()
}

func (this *MyQueue) Peek() int {
	if this.right.empty() {
		for !this.left.empty() {
			this.right.push(this.left.pop())
		}
	}
	return this.right.peek()
}

func (this *MyQueue) Empty() bool {
	return this.left.empty() && this.right.empty()
}

func (this *MyQueue) String() string {
	return fmt.Sprintf("left:%v\nright:%v\n\n", this.left, this.right)
}
