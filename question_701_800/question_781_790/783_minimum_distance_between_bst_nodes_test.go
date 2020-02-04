package question_781_790

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_minDiffInBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNode(4, NewNode(2, NewNode(1), NewNode(3)), NewNode(6)), 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDiffInBST(tt.root); got != tt.want {
				t.Errorf("minDiffInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
