package question_101_110

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sortedListToBST(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *TreeNode
	}{
		{nil, nil},
		{NewList(1), NewNodeV2(1)},
		{NewList(1, 2), NewNodeV2(2, 1)},
		{NewList(1, 2, 3), NewNodeV2(2, 1, 3)},
		{NewList(-10, -3, 0, 5, 9), NewNodeV2(0, -3, 9, -10, math.MaxInt32, 5)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortedListToBST(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
