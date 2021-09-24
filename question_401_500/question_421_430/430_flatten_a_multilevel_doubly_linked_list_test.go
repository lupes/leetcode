package question_421_430

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	node1 := &Node3{Val: 1}
	node2 := &Node3{Val: 2, Prev: node1}
	node3 := &Node3{Val: 3, Prev: node2}
	node4 := &Node3{Val: 4, Prev: node3}
	node5 := &Node3{Val: 5, Prev: node4}
	node6 := &Node3{Val: 6, Prev: node5}
	node7 := &Node3{Val: 7}
	node8 := &Node3{Val: 8, Prev: node7}
	node9 := &Node3{Val: 9, Prev: node8}
	node10 := &Node3{Val: 10, Prev: node9}
	node11 := &Node3{Val: 11}
	node12 := &Node3{Val: 12, Prev: node11}

	node1.Next, node2.Next, node3.Next, node4.Next, node5.Next = node2, node3, node4, node5, node6
	node7.Next, node8.Next, node9.Next = node8, node9, node10
	node11.Next = node12

	node3.Child = node7
	node8.Child = node11

	got1 := &Node3{Val: 1}
	got2 := &Node3{Val: 2, Prev: got1}
	got3 := &Node3{Val: 3, Prev: got2}
	got7 := &Node3{Val: 7, Prev: got3}
	got8 := &Node3{Val: 8, Prev: got7}
	got11 := &Node3{Val: 11, Prev: got8}
	got12 := &Node3{Val: 12, Prev: got11}
	got9 := &Node3{Val: 9, Prev: got12}
	got10 := &Node3{Val: 10, Prev: got9}
	got4 := &Node3{Val: 4, Prev: got10}
	got5 := &Node3{Val: 5, Prev: got4}
	got6 := &Node3{Val: 6, Prev: got5}

	got1.Next, got2.Next, got3.Next, got7.Next, got8.Next, got11.Next, got12.Next, got9.Next, got10.Next, got4.Next, got5.Next =
		got2, got3, got7, got8, got11, got12, got9, got10, got4, got5, got6

	tests := []struct {
		root *Node3
		want *Node3
	}{
		{node1, got1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := flatten(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
