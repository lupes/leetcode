package question_531_540

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_convertBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{NewNodeV2(5, 2, 13), NewNodeV2(18, 20, 13)},
		{NewNodeV2(5, 2, 13, 1, 3, 6, 17), NewNodeV2(41, 46, 30, 47, 44, 36, 17)},
		{
			NewNodeV2(0, -1, 2, -3, math.MaxInt32, math.MaxInt32, 4),
			NewNodeV2(6, 5, 6, 2, math.MaxInt32, math.MaxInt32, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := convertBST(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
