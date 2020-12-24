package question_891_900

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_increasingBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{
			NewNodeV2(5, 3, 6, 2, 4, Null, 8, 1, Null, Null, Null, 7, 9),
			NewNodeV2(1, Null, 2, Null, 3, Null, 4, Null, 5, Null, 6, Null, 7, Null, 8, Null, 9),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := increasingBST(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
