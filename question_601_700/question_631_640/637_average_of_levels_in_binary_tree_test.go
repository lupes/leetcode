package question_631_640

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_averageOfLevels(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []float64
	}{
		{NewNodeV2(3, 9, 20, math.MaxInt32, math.MaxInt32, 15, 7), []float64{3, 14.5, 11}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := averageOfLevels(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
