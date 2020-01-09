package question_691_700

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_searchBST(t *testing.T) {
	root := NewNode(4, NewNode(2, NewNode(1), NewNode(3)), NewNode(7))
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{nil, 2, nil},
		{root, 8, nil},
		{root, 2, root.Left},
		{root, 3, root.Left.Right},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := searchBST(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
