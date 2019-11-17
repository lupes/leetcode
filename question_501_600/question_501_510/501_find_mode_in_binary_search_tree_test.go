package question_501_510

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_findMode(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		//{nil, nil},
		//{&TreeNode{Val: 0}, []int{0}},
		{&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, []int{1, 2, 3}},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMode(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
