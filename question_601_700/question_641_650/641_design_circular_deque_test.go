package question_641_650

import (
	"fmt"
	"testing"
)

func TestMyCircularDeque(t *testing.T) {

	got := Constructor(3)

	fmt.Printf("got:%v want:true\n", got.InsertLast(1))
	fmt.Printf("got:%v want:true\n", got.InsertLast(2))
	fmt.Printf("got:%v want:true\n", got.InsertFront(3))
	fmt.Printf("got:%v want:false\n", got.InsertFront(4))

	fmt.Printf("got:%v want:2\n", got.GetRear())         // 返回 2
	fmt.Printf("got:%v want:true\n", got.IsFull())       // 返回 true
	fmt.Printf("got:%v want:true\n", got.DeleteLast())   // 返回 true
	fmt.Printf("got:%v want:true\n", got.InsertFront(4)) // 返回 true
	fmt.Printf("got:%v want:4\n", got.GetFront())        // 返回4

}
