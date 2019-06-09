package question_91_100

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	var root1 *TreeNode
	var root2 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	var root3 = &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"test", root1, nil},
		{"test", root2, []int{2, 3, 1}},
		{"test", root3, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
