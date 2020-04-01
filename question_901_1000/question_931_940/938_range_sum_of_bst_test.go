package question_931_940

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_rangeSumBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		L    int
		R    int
		want int
	}{
		{NewNode(10, NewNode(5, NewNode(3), NewNode(7)), NewNode(15, nil, NewNode(18))), 7, 15, 32},
		{NewNode(10, NewNode(5, NewNode(3, NewNode(1)), NewNode(7, NewNode(6))), NewNode(15, NewNode(13), NewNode(18))), 6, 10, 23},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rangeSumBST(tt.root, tt.L, tt.R); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
