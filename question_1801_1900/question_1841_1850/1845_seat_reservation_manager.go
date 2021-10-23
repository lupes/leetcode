package question_1841_1850

import (
	"container/heap"
)

// 1845. 座位预约管理系统
// https://leetcode-cn.com/problems/seat-reservation-manager/
// 设计 堆（优先队列）

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *IntHeap) Pop() interface{} {
	x := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return x
}

type SeatManager struct {
	seats *IntHeap
}

func Constructor(n int) SeatManager {
	var h = &IntHeap{}
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{
		seats: h,
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.seats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.seats, seatNumber)
}
