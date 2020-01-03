package question_651_660

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		nums []int
		want *TreeNode
	}{
		{[]int{1}, NewNode(1)},
		{[]int{-1}, NewNode(-1)},
		{[]int{-1, 2}, NewNode(2, NewNode(-1))},
		{[]int{-1, 2, 1}, NewNode(2, NewNode(-1), NewNode(1))},
		{[]int{3, 2, 1, 6, 0, 5}, NewNode(6, NewNode(3, nil, NewNode(2, nil, NewNode(1))), NewNode(5, NewNode(0)))},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
