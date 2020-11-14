package question_1361_1370

// 1381. 设计一个支持增量操作的栈
// https://leetcode-cn.com/problems/design-a-stack-with-increment-operation/
// Topics: 栈 设计

type CustomStack struct {
	len    int
	size   int
	values []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		len:    0,
		size:   maxSize,
		values: make([]int, 0, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if this.len == this.size {
		return
	}
	this.values = append(this.values, x)
	this.len++
}

func (this *CustomStack) Pop() int {
	if this.len == 0 {
		return -1
	}

	value := this.values[this.len-1]
	this.values = this.values[:this.len-1]
	this.len--
	return value
}

func (this *CustomStack) Increment(k int, val int) {
	if this.len < k {
		k = this.len
	}
	for i := 0; i < k; i++ {
		this.values[i] += val
	}
}
