package question_111_120

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"test", nil, nil},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			n1 := tt.root
			n2 := tt.want
			for n1 != nil && n2 != nil {
				fmt.Printf("%d %d\n", n1.Val, n2.Val)
				n1 = n1.Right
				n2 = n2.Right
			}
		})
	}
}
