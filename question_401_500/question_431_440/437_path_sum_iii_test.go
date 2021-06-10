package question_431_440

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_pathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		sum  int
		want int
	}{
		{NewNodeV2(10, 5, -3, 3, 2, Null, 11, 3, -2, Null, 1), 8, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pathSum(tt.root, tt.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
