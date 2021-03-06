package question_111_120

import "testing"

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"test", nil, 0},
		{"test", &TreeNode{Val: 1}, 1},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, 2},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
