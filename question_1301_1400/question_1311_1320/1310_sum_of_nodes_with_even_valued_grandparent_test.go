package question_1311_1320

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sumEvenGrandparent(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(6, 7, 8, 2, 7, 1, 3, 9, Null, 1, 4, Null, Null, Null, 5), 18},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumEvenGrandparent(tt.root); got != tt.want {
				t.Errorf("sumEvenGrandparent() = %v, want %v", got, tt.want)
			}
		})
	}
}
