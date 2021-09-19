package offer

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_rightSideView(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{NewNodeV2(1, 2, 3, Null, 5, Null, 4), []int{1, 3, 4}},
		{NewNodeV2(1, Null, 3), []int{1, 3}},
		{nil, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
