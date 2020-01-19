package question_701_710

import (
	"fmt"
	"testing"
)

//func TestMyLinkedList_FindIndex(t *testing.T) {
//	head := Constructor5()
//	head.AddAtHead(4)
//	head.AddAtHead(3)
//	head.AddAtHead(2)
//	head.AddAtHead(1)
//	tests := []struct {
//		index int
//		want  int
//	}{
//		{3, 1},
//	}
//	for _, tt := range tests {
//		t.Run("test", func(t *testing.T) {
//			//got := head.findIndex(tt.index)
//			//fmt.Printf("%s\n", got)
//		})
//	}
//}

func TestConstructor_1(t *testing.T) {
	this := Constructor5()
	this.AddAtHead(1)
	fmt.Printf("%+v\n", this)
	this.AddAtTail(3)
	fmt.Printf("%+v\n", this)
	this.AddAtIndex(1, 2)
	fmt.Printf("%+v\n", this)

	fmt.Printf("%d\n", this.Get(1))

	this.DeleteAtIndex(1)
	fmt.Printf("%+v\n", this)

	fmt.Printf("%d\n", this.Get(1))
}

func TestConstructor_2(t *testing.T) {

	// ["addAtHead","addAtHead","addAtHead", [7],[2],[1],
	this := Constructor5()
	this.AddAtHead(7)
	fmt.Printf("%+v\n", this)
	this.AddAtHead(2)
	fmt.Printf("%+v\n", this)
	this.AddAtHead(1)
	fmt.Printf("%+v\n", this)

	// addAtIndex [3,0],
	this.AddAtIndex(3, 0)
	fmt.Printf("%+v\n", this)

	// deleteAtIndex [2],
	this.DeleteAtIndex(2)
	fmt.Printf("%+v\n", this)

	// addAtHead [6],
	this.AddAtHead(6)
	fmt.Printf("%+v\n", this)

	// addAtTail [4],
	this.AddAtTail(4)
	fmt.Printf("%+v\n", this)

	// get [4],
	fmt.Printf("%d\n", this.Get(4))

	// addAtHead [4],
	this.AddAtHead(4)
	fmt.Printf("%+v\n", this)

	// addAtIndex [5,0]
	this.AddAtIndex(5, 0)
	fmt.Printf("%+v\n", this)

	// addAtHead [6]
	this.AddAtHead(6)
	fmt.Printf("%+v\n", this)
}

func TestConstructor_3(t *testing.T) {
	this := Constructor5()
	this.AddAtHead(1)
	fmt.Printf("%+v\n", this)
	this.DeleteAtIndex(1)
	fmt.Printf("%+v\n", this)
	this.DeleteAtIndex(0)
	fmt.Printf("%+v\n", this)
	this.AddAtTail(3)
	fmt.Printf("%+v\n", this)
	this.AddAtIndex(1, 2)
	fmt.Printf("%+v\n", this)

	fmt.Printf("%d\n", this.Get(1))

	this.DeleteAtIndex(0)
	fmt.Printf("%+v\n", this)

	fmt.Printf("%d\n", this.Get(0))
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	head := Constructor5()
	head.AddAtHead(4)
	head.AddAtHead(3)
	head.AddAtHead(2)
	head.AddAtHead(1)

	fmt.Printf("%+v\n", head)
}
