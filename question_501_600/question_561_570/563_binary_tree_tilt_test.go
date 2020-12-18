package question_561_570

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_findTiltHelper(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(4, 2, 9, 3, 5, Null, 7), 15},
		{NewNodeV2(21, 7, 14, 1, 1, 2, 2, 3, 3), 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findTilt(tt.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
