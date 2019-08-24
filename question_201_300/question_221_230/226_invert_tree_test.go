package question_221_230

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"test", nil, nil},
		{"test", &TreeNode{Val: 4}, &TreeNode{Val: 4}},
		{"test", &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 4, Right: &TreeNode{Val: 2}}},
		{"test", &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
		}, &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
