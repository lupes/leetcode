package question_121_130

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sumNumbers(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(1, 2, 3), 25},
		{NewNodeV2(1, 2, 3, 4, 5, 6, 7), 25},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumNumbers(tt.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
