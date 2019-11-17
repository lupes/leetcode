package question_111_120

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		sum  int
		want [][]int
	}{
		{"test", nil, 0, nil},
		{"test", &TreeNode{Val: 0}, 0, [][]int{{0}}},
		{"test", &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3}}, 2, nil},
		{"test", &TreeNode{Val: 0, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 3}}, 3, [][]int{{0, 1, 2}, {0, 3}}},
		{"test", &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}, 22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.root, tt.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
