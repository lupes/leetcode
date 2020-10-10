package question_281_290

// 284. 顶端迭代器
// https://leetcode-cn.com/problems/peeking-iterator
// Topics: 设计

type Iterator struct {
	arr []int
}

func (this *Iterator) hasNext() bool {
	return len(this.arr) != 0
}
func (this *Iterator) next() int {
	tmp := this.arr[0]
	this.arr = this.arr[1:]
	return tmp
}

type PeekingIterator struct {
	arr []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	this := &PeekingIterator{}
	for iter.hasNext() {
		this.arr = append(this.arr, iter.next())
	}
	return this
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.arr) != 0
}

func (this *PeekingIterator) next() int {
	tmp := this.arr[0]
	this.arr = this.arr[1:]
	return tmp
}

func (this *PeekingIterator) peek() int {
	return this.arr[0]
}
