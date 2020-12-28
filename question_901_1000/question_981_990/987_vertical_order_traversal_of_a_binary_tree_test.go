package question_981_990

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_verticalTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{NewNodeV2(3, 9, 20, Null, Null, 15, 7), [][]int{{9}, {3, 15}, {20}, {7}}},
		{NewNodeV2(1, 2, 3, 4, 5, 6, 7), [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
		{NewNodeV2(0, 5, 1, 9, Null, 2, Null, Null, Null, Null, 3, 4, 8, 6, Null, Null, Null, 7), [][]int{{9, 7}, {5, 6}, {0, 2, 4}, {1, 3}, {8}}},
		{NewNodeV2(0, 8, 1, Null, Null, 3, 2, Null, 4, 5, Null, Null, 7, 6), [][]int{{8}, {0, 3, 6}, {1, 4, 5}, {2, 7}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := verticalTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalTraversal() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
