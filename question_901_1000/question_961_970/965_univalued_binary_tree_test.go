package question_961_970

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isUnivalTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{NewNode(1, NewNode(1), NewNode(1)), true},
		{NewNode(2, NewNode(1), NewNode(1)), false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isUnivalTree(tt.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
