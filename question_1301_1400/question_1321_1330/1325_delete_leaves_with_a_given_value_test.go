package question_1321_1330

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_removeLeafNodes(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		target int
		want   *TreeNode
	}{
		{
			NewNodeV2(1, 2, 3, 2, Null, 2, 4), 2, NewNodeV2(1, Null, 3, Null, 4),
		},
		{
			NewNodeV2(1, 3, 3, 3, 2), 3, NewNodeV2(1, 3, Null, Null, 2),
		},
		{
			NewNodeV2(1, 2, Null, 2, Null, 2), 2, NewNodeV2(1),
		},
		{
			NewNodeV2(1, 1, 1), 1, nil,
		},
		{
			NewNodeV2(1, 2, 3), 1, NewNodeV2(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeLeafNodes(tt.root, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
