package question_141_150

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_postorderTraversal(t *testing.T) {
	//test1 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	//test2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{nil, nil},
		{NewNode(1, nil, NewNode(2, NewNode(3))), []int{3, 2, 1}},
		{NewNode(1, NewNode(2, nil, NewNode(4, NewNode(6))), NewNode(3, nil, NewNode(5))), []int{6, 4, 2, 5, 3, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := postorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
