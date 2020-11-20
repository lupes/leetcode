package question_541_550

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(1), 0},
		{NewNodeV2(1, 1), 1},
		{NewNodeV2(1, 1, 1), 2},
		{NewNodeV2(1, 2, 3, 4, 5), 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
