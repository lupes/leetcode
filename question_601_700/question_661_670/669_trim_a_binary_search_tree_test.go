package question_661_670

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_trimBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		L    int
		R    int
		want *TreeNode
	}{
		{NewNodeV2(1, 0, 2), 1, 2, NewNodeV2(1, math.MaxInt32, 2)},
		{NewNodeV2(3, 0, 4, math.MaxInt32, 2, math.MaxInt32, math.MaxInt32, 1), 1, 3,
			NewNodeV2(3, 2, math.MaxInt32, 1)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := trimBST(tt.root, tt.L, tt.R); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
