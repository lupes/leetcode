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

func TestConstructor5(t *testing.T) {
	tests := []struct {
		funcs []string
		args  [][]int
	}{
		{
			funcs: []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"},
			args:  [][]int{{2}, {1}, {2}, {7}, {3}, {2}, {5}, {5}, {5}, {6}, {4}},
		},
		{
			funcs: []string{"MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"},
			args:  [][]int{{1}, {3}, {1, 2}, {1}, {0}, {0}},
		},
		{
			funcs: []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"},
			args:  [][]int{{2}, {1}, {2}, {7}, {3}, {2}, {5}, {5}, {5}, {6}, {4}},
		},
		{
			funcs: []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"},
			args:  [][]int{{2}, {1}, {2}, {7}, {3}, {2}, {5}, {5}, {5}, {6}, {4}},
		},
		{
			funcs: []string{"MyLinkedList", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "get", "deleteAtIndex", "deleteAtIndex"},
			args:  [][]int{{2}, {1}, {2}, {7}, {3}, {2}, {5}, {5}, {5}, {6}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			this := Constructor5()
			for i, fun := range tt.funcs[1:] {
				switch fun {
				case "addAtHead":
					this.AddAtHead(tt.args[i][0])
				case "addAtTail":
					this.AddAtTail(tt.args[i][0])
				case "addAtIndex":
					this.AddAtIndex(tt.args[i][0], tt.args[i][1])
				case "deleteAtIndex":
					this.DeleteAtIndex(tt.args[i][0])
				case "get":
					fmt.Printf("%d: Get:%d\n", i, this.Get(tt.args[i][0]))
				}
				fmt.Printf("%d: %s %+v %+v\n", i, fun, tt.args[i], this)
			}
		})
	}
}
