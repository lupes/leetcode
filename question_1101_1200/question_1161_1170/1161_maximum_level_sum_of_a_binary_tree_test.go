package question_1161_1170

import (
	"math"
	"testing"

	. "github.com/lupes/leetcode/common"
)

// 989
// null 10250
//      98693 -89388
//                -32127

func Test_maxLevelSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNodeV2(1, 7, 0, -8, math.MaxInt32, math.MaxInt32), 2},
		{NewNodeV2(989, math.MaxInt32, 10250, 98693, -89388, math.MaxInt32, math.MaxInt32, math.MaxInt32, -32127), 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxLevelSum(tt.root); got != tt.want {
				t.Errorf("maxLevelSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
