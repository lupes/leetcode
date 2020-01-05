package question_661_670

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_widthOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{nil, 0},
		{NewNode(1), 1},
		{NewNode(1, NewNode(3)), 1},
		{NewNode(1, NewNode(3, NewNode(5), NewNode(3)), NewNode(2, nil, NewNode(9))), 4},
		{NewNode(1, NewNode(3, NewNode(5), NewNode(3)), NewNode(2, NewNode(9))), 3},
		{NewNode(1, NewNode(3, NewNode(5), NewNode(3))), 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := widthOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
