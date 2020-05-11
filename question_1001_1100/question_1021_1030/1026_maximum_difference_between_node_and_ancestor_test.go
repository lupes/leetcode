package question_1021_1030

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_maxAncestorDiff(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNode(1, NewNode(2)), 1},
		{NewNode(8,
			NewNode(3, NewNode(1), NewNode(6, NewNode(4), NewNode(7))),
			NewNode(10, nil, NewNode(14, NewNode(13)))), 7},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxAncestorDiff(tt.root); got != tt.want {
				t.Errorf("maxAncestorDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
