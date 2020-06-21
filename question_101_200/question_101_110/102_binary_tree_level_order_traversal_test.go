package question_101_110

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{"test", nil, nil},
		{"test", &TreeNode{Val: 1}, [][]int{{1}}},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, [][]int{{1}, {2, 3}}},
		{"test", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, [][]int{{1}, {2, 3}, {4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
