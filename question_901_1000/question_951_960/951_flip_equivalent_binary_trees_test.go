package question_951_960

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_flipEquiv(t *testing.T) {
	tests := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{nil, nil, true},
		{NewNode(1), NewNode(1), true},
		{NewNode(1, NewNode(2)), NewNode(1, NewNode(2)), true},
		{NewNode(1, nil, NewNode(2)), NewNode(1, NewNode(2)), true},
		{NewNode(1, nil, NewNode(3)), NewNode(1, NewNode(2)), false},
		{NewNode(1, NewNode(2, NewNode(4), NewNode(5, NewNode(7), NewNode(8))), NewNode(3, NewNode(6))),
			NewNode(1, NewNode(3, nil, NewNode(6)), NewNode(2, NewNode(4), NewNode(5, NewNode(8), NewNode(7)))), true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := flipEquiv(tt.root1, tt.root2); got != tt.want {
				t.Errorf("flipEquiv() = %v, want %v", got, tt.want)
			}
		})
	}
}
