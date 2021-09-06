package question_1661_1670

import (
	"fmt"
	"testing"
)

//["FrontMiddleBackQueue","","","","","","","popMiddle","popMiddle","pushMiddle","pushMiddle","popFront"]
//[[],[],[],[],[],[],[],[],[],[],[],[]]

func TestFrontMiddleBackQueue(t *testing.T) {
	queue := Constructor()

	queue.PushFront(1)
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PushBack(2)
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PushMiddle(3)
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PushMiddle(4)
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PopFront()
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PopMiddle()
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PopMiddle()
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PopBack()
	fmt.Printf("%s\n", queue.Head.Next)
	queue.PopFront()
	fmt.Printf("%s\n", queue.Head.Next)

}
