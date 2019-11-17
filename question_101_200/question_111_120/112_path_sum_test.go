package question_111_120

import "testing"

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		sum  int
		want bool
	}{
		{"test", nil, 1, false},
		{"test", nil, 0, false},
		{"test", &TreeNode{Val: 1}, 0, false},
		{"test", &TreeNode{Val: 1}, 1, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5}}, 7, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
