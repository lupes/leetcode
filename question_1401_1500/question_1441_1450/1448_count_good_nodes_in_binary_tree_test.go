package question_1441_1450

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_goodNodes(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(3, 1, 4, 3, Null, 1, 5), 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := goodNodes(tt.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
