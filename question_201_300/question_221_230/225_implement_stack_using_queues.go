package question_221_230

// 225. 用队列实现栈
// https://leetcode-cn.com/problems/implement-stack-using-queues/
// Topics: 栈 设计

type MyStack struct {
	value int
	next  *MyStack
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	n := &MyStack{
		value: x,
		next:  this.next,
	}
	this.next = n
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	t := this.next
	this.next = t.next
	return t.value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.next.value
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.next == nil
}
