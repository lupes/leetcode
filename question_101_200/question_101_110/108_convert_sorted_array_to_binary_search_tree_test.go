package question_101_110

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{"test", nil, nil},
		{"test", []int{1}, &TreeNode{Val: 1}},
		{"test", []int{1, 2}, &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
		{"test", []int{1, 2, 3}, &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
