package question_141_150

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	test1 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	test2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test#1", args{test1}, []int{3, 2, 1}},
		{"test#2", args{test2}, []int{6, 4, 2, 5, 3, 1}},
		{"test#3", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
