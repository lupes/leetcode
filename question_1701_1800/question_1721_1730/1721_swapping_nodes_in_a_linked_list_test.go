package question_1721_1730

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_swapNodes(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{NewList(1, 2, 3, 4, 5), 1, NewList(5, 2, 3, 4, 1)},
		{NewList(1, 2, 3, 4, 5), 2, NewList(1, 4, 3, 2, 5)},
		{NewList(1, 2, 3, 4, 5), 3, NewList(1, 2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 4, NewList(1, 4, 3, 2, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(5, 2, 3, 4, 1)},
		{NewList(7, 9, 6, 6, 7, 8, 3, 0, 9, 5), 5, NewList(7, 9, 6, 6, 8, 7, 3, 0, 9, 5)},
		{NewList(1), 1, NewList(1)},
		{NewList(1, 2), 1, NewList(2, 1)},
		{NewList(1, 2, 3), 2, NewList(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := swapNodes(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
