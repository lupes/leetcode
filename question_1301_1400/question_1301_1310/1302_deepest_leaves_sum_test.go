package question_1301_1310

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_deepestLeavesSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(1, 2, 3, 4, 5, Null, 6, 7, Null, Null, Null, Null, 8), 15},
		{NewNodeV2(1, 2, 3, 4, 5, Null, 6), 15},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := deepestLeavesSum(tt.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
