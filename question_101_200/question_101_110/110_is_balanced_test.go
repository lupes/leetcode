package question_101_110

import "testing"

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"test", nil, true},
		{"test", &TreeNode{Val: 1}, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
