package question_951_960

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isCompleteTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		//{nil, false},
		//{NewNode(1), true},
		//{NewNode(1, NewNode(2)), true},
		//{NewNode(1, NewNode(2), NewNode(3)), true},
		//{NewNode(1, nil, NewNode(3)), false},
		//{NewNode(1, NewNode(2, NewNode(4), NewNode(5)), NewNode(3)), true},
		{NewNode(1, NewNode(2, NewNode(4), NewNode(5)), NewNode(3, NewNode(6))), true},
		{NewNode(1, NewNode(2, NewNode(4), NewNode(5)), NewNode(3, nil, NewNode(6))), false},
		{NewNode(1,
			NewNode(2, NewNode(4, NewNode(8, NewNode(15)), NewNode(9)), NewNode(5, NewNode(10), NewNode(11))),
			NewNode(3, NewNode(6, NewNode(12), NewNode(13)), NewNode(7))), false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isCompleteTree(tt.root); got != tt.want {
				t.Errorf("isCompleteTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
