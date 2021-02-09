package question_1451_1460

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(9), 1},
		{NewNodeV2(2, 2), 1},
		{NewNodeV2(2, 2, 2), 2},
		{NewNodeV2(2, 3, 1, 3, 1, Null, 1), 2},
		{NewNodeV2(2, 1, 1, 1, 3, Null, Null, Null, Null, Null, 1), 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
