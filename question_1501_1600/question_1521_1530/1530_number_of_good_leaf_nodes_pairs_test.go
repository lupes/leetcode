package question_1521_1530

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_countPairs(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		distance int
		want     int
	}{
		{NewNodeV2(1, 2, 3, Null, 4), 3, 1},
		{NewNodeV2(100), 1, 0},
		{NewNodeV2(1, 1, 1), 2, 1},
		{NewNodeV2(1, 2, 3, 4, 5, 6, 7), 3, 2},
		{NewNodeV2(7, 1, 4, 6, Null, 5, 3, Null, Null, Null, Null, Null, 2), 3, 1},
		{NewNodeV2(15, 66, 55, 97, 60, 12, 56, Null, 54, Null, 49, Null, 9, Null, Null, Null, Null, Null, 90), 5, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countPairs(tt.root, tt.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
