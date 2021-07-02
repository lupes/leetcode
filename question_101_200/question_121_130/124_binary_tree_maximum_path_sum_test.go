package question_121_130

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(1, 2, 3), 6},
		{NewNodeV2(1, 2), 3},
		{NewNodeV2(-3), -3},
		{NewNodeV2(1, -2, 3), 4},
		{NewNodeV2(-1, 2, 3), 4},
		{NewNodeV2(-10, 9, 20, Null, Null, 15, 7), 42},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
