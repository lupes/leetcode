package question_1071_1080

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sufficientSubset(t *testing.T) {
	tests := []struct {
		root  *TreeNode
		limit int
		want  *TreeNode
	}{
		{NewNodeV2(1, math.MaxInt32, 3, -99, 7, 12, 13, -99, 14), 1,
			NewNodeV2(1, math.MaxInt32, 3, math.MaxInt32, math.MaxInt32, 7, math.MaxInt32, 14),
		},
		{NewNodeV2(1, 2, 3, 4, -99, -99, 7, 8, 9, -99, -99, 12, 13, -99, 14), 1,
			NewNodeV2(1, 2, 3, 4, math.MaxInt32, math.MaxInt32, 7, 8, 9, math.MaxInt32, 14),
		},
		//{
		//	NewNodeV2(5, 4, 8, 11, math.MaxInt32, 17, 4, 7, 1, math.MaxInt32, math.MaxInt32, 5, 3), 22,
		//	NewNodeV2(5, 4, 8, 11, math.MaxInt32, 17, 4, 7, math.MaxInt32, math.MaxInt32, math.MaxInt32, 5),
		//},
		//{
		//	NewNodeV2(5, -6, -6), 0,
		//	nil,
		//},
		//{NewNodeV2(1, 2, -3, -5, math.MaxInt32, 4, math.MaxInt32), -1, NewNodeV2(1, math.MaxInt32, -3, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sufficientSubset(tt.root, tt.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sufficientSubset() = \n got  %v,\n want %v", got, tt.want)
			}
		})
	}
}
