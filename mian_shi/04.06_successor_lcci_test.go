package mian_shi

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_inorderSuccessor(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		want *TreeNode
	}{
		{NewNodeV2(2, 1, 3), NewNode(1), NewNode(2)},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), NewNode(2), NewNode(3)},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), NewNode(3), NewNode(4)},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), NewNode(4), NewNode(5)},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), NewNode(5), NewNode(6)},
		{NewNodeV2(5, 3, 6, 2, 4, Null, Null, 1), NewNode(6), nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := inorderSuccessor(tt.root, tt.p)
			if !((got == nil && tt.want == nil) || (got != nil && tt.want != nil && got.Val == tt.want.Val)) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
