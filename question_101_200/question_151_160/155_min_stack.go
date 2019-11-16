package question_151_160

type MinStack struct {
	stacks []int
	sorts  []int
	size   int
}

// 155. 最小栈
// https://leetcode-cn.com/problems/min-stack

func Constructor() MinStack {
	return MinStack{
		stacks: make([]int, 0),
		sorts:  make([]int, 0),
		size:   0,
	}
}

func (this *MinStack) Push(x int) {
	this.stacks = append(this.stacks, x)
	for i, n := range this.sorts {
		if n >= x {
			var t = make([]int, len(this.sorts[i:]))
			copy(t, this.sorts[i:])
			this.sorts = append(this.sorts[:i], x)
			this.sorts = append(this.sorts, t...)
			goto Next
		}
	}
	this.sorts = append(this.sorts, x)
Next:
	this.size++
}

func (this *MinStack) Pop() {
	x := this.stacks[this.size-1]
	this.stacks = this.stacks[:this.size-1]
	for i, n := range this.sorts {
		if n == x {
			if i == this.size-1 {
				this.sorts = this.sorts[:this.size-1]
			} else {
				this.sorts = append(this.sorts[:i], this.sorts[i+1:]...)
			}
		}
	}
	this.size--
}

func (this *MinStack) Top() int {
	return this.stacks[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.sorts[0]
}
