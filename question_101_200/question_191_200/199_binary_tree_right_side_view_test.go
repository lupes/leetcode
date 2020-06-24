package question_191_200

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{NewNodeV2(1, 2, 3, math.MaxInt32, 5, math.MaxInt32, 4), []int{1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
