package question_231_240

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want int
	}{
		{NewNodeV2(6, 2, 8, 0, 4, 7, 9, Null, Null, 3, 5), NewNodeV2(2), NewNodeV2(8), 6},
		{NewNodeV2(6, 2, 8, 0, 4, 7, 9, Null, Null, 3, 5), NewNodeV2(4), NewNodeV2(3), 2},
		{NewNodeV2(2, 3, 1), NewNodeV2(1), NewNodeV2(3), 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lowestCommonAncestor(tt.root, tt.p, tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
