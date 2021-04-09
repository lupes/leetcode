package question_1601_1610

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isEvenOddTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{NewNodeV2(1, 10, 4, 3, Null, 7, 9, 12, 8, 6, Null, Null, 2), true},
		{NewNodeV2(5, 4, 2, 3, 3, 7), false},
		{NewNodeV2(5, 9, 1, 3, 5, 7), false},
		{NewNodeV2(1), true},
		{NewNodeV2(11, 8, 6, 1, 3, 9, 11, 30, 20, 18, 16, 12, 10, 4, 2, 17), true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isEvenOddTree(tt.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
