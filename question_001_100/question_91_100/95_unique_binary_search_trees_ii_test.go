package question_91_100

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*TreeNode
	}{
		{1, []*TreeNode{{Val: 1}}},
		{2, []*TreeNode{{Val: 1, Right: &TreeNode{Val: 2}}, {Val: 2, Left: &TreeNode{Val: 1}}}},
		{3, []*TreeNode{{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, {Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}, {Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, {Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}}, {Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := generateTrees(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %+v, want %v", got, tt.want)
			}
		})
	}
}
