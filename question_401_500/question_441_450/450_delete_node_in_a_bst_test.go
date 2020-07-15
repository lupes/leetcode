package question_441_450

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		root *TreeNode
		key  int
		want *TreeNode
	}{
		{NewNodeV2(5, 3, 6, 2, 4, math.MaxInt32, 7), 3, NewNodeV2(5, 2, 6, math.MaxInt32, 4, math.MaxInt32, 7)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := deleteNode(tt.root, tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
