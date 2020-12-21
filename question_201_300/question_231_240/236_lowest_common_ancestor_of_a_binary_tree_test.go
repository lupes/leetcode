package question_231_240

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_lowestCommonAncestor2(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{NewNodeV2(3, 5, 1, 6, 2, 0, 8, Null, Null, 7, 4), &TreeNode{Val: 5}, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		{NewNodeV2(3, 5, 1, 6, 2, 0, 8, Null, Null, 7, 4), &TreeNode{Val: 5}, &TreeNode{Val: 4}, &TreeNode{Val: 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lowestCommonAncestor2(tt.root, tt.p, tt.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got, tt.want)
			}
		})
	}
}
