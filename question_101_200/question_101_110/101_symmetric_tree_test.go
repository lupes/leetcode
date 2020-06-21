package question_101_110

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"test", nil, true},
		{"test", &TreeNode{Val: 1}, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, false},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, true},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
