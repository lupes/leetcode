package question_701_710

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_insertIntoBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{nil, 1, NewNode(1)},
		{NewNode(4, NewNode(2, NewNode(1), NewNode(3)), NewNode(7)), 5, NewNode(4, NewNode(2, NewNode(1), NewNode(3)), NewNode(7, NewNode(5)))},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := insertIntoBST(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
