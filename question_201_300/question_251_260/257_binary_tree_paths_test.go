package question_251_260

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []string
	}{
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, []string{"1->2->5", "1->3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
