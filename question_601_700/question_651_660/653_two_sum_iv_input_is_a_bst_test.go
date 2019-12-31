package question_651_660

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_findTarget(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want bool
	}{
		{NewNode(5, NewNode(3, NewNode(2), NewNode(4)), NewNode(6, nil, NewNode(7))), 9, true},
		{NewNode(5, NewNode(3, NewNode(2), NewNode(4)), NewNode(6, nil, NewNode(7))), 27, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findTarget(tt.root, tt.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
