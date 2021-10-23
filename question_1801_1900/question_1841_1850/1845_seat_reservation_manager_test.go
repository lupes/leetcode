package question_1841_1850

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestSeatManager(t *testing.T) {
	sm := Constructor(5)
	fmt.Printf("%d\n", sm.Reserve())
	fmt.Printf("%d\n", sm.Reserve())
	sm.Unreserve(2)
	fmt.Printf("%d\n", sm.Reserve())
	fmt.Printf("%d\n", sm.Reserve())
	fmt.Printf("%d\n", sm.Reserve())
	fmt.Printf("%d\n", sm.Reserve())
	sm.Unreserve(5)
}

func TestIntHeap(t *testing.T) {
	h := &IntHeap{}
	heap.Push(h, 2)
	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 9)
	heap.Push(h, 8)
	heap.Push(h, -1)
	heap.Push(h, 100)
	heap.Push(h, -2)

	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
	fmt.Printf("%v\n", heap.Pop(h))
}
