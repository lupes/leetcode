package question_111_120

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	var node1 = &Node{Val: 1}
	var node2 = &Node{Val: 2}
	var node3 = &Node{Val: 3}
	var node4 = &Node{Val: 4}
	var node5 = &Node{Val: 5}
	var node6 = &Node{Val: 6}
	var node7 = &Node{Val: 7}
	node1.Left = node2
	node1.Right = node3
	node2.Next = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	tests := []struct {
		root *Node
		want *Node
	}{
		{
			root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val:   4,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: &Node{
						Val:   5,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Next: nil,
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val:   6,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: &Node{
						Val:   7,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Next: nil,
				},
			},
			want: node1,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := connect(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
