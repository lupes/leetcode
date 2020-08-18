package question_991_1000

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_insertIntoMaxTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{NewNodeV2(4, 1, 3, Null, Null, 2), 5, NewNodeV2(5, 4, Null, 1, 3, Null, Null, 2)},
		{NewNodeV2(5, 2, 4, Null, 1), 3, NewNodeV2(5, 2, 4, Null, 1, Null, 3)},
		{NewNodeV2(5, 2, 3, Null, 1), 4, NewNodeV2(5, 2, 4, Null, 1, 3)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := insertIntoMaxTree(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoMaxTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
