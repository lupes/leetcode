package question_1021_1030

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sumRootToLeaf(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNode(1, NewNode(0, NewNode(0), NewNode(1)), NewNode(1, NewNode(0), NewNode(1))), 22},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumRootToLeaf(tt.root); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
