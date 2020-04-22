package question_991_1000

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isCousins(t *testing.T) {
	tests := []struct {
		root *TreeNode
		x    int
		y    int
		want bool
	}{
		{NewNode(1, NewNode(2, NewNode(4)), NewNode(3)), 4, 3, false},
		{NewNode(1, NewNode(2, NewNode(4)), NewNode(3)), 1, 2, false},
		{NewNode(1, NewNode(2, NewNode(4)), NewNode(3)), 2, 3, false},
		{NewNode(1, NewNode(2, nil, NewNode(4)), NewNode(3, nil, NewNode(5))), 4, 5, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isCousins(tt.root, tt.x, tt.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
