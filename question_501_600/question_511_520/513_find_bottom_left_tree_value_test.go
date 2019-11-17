package qustion_511_520

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"test", &TreeNode{Val: 1}, 1},
		{"test", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, 1},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, 4},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
