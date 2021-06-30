package question_221_230

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want int
	}{
		{NewNodeV2(3, 1, 4, Null, 2), 1, 1},
		{NewNodeV2(3, 1, 4, Null, 2), 2, 2},
		{NewNodeV2(3, 1, 4, Null, 2), 3, 3},
		{NewNodeV2(3, 1, 4, Null, 2), 4, 4},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), 3, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kthSmallest(tt.root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
