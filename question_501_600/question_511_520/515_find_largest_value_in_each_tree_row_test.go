package qustion_511_520

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"test", &TreeNode{Val: 1}, []int{1}},
		{"test", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, []int{2, 3}},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, []int{1, 3, 4}},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}}}, []int{1, 3, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
