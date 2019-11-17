package question_91_100

import "testing"

func Test_isSameTree(t *testing.T) {
	var (
		p1 *TreeNode
		q1 *TreeNode
		p2 = &TreeNode{Val: 1}
		q2 = &TreeNode{Val: 1}
		p3 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
		q3 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
		p4 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
		q4 = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	)
	tests := []struct {
		name string
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		{"test", p1, q1, true},
		{"test", p2, q2, true},
		{"test", p1, q2, false},
		{"test", p2, q1, false},
		{"test", p3, q3, true},
		{"test", p3, q2, false},
		{"test", p4, q4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.p, tt.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
