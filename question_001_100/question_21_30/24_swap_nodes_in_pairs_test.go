package question_0011_0020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{nil, nil},
		{NewList(1), NewList(1)},
		{NewList(1, 2), NewList(2, 1)},
		{NewList(1, 2, 3), NewList(2, 1, 3)},
		{NewList(1, 2, 3, 4), NewList(2, 1, 4, 3)},
		{NewList(1, 2, 3, 4, 5), NewList(2, 1, 4, 3, 5)},
		{NewList(1, 2, 3, 4, 5, 6), NewList(2, 1, 4, 3, 6, 5)},
		{NewList(1, 2, 3, 4, 5, 6, 7), NewList(2, 1, 4, 3, 6, 5, 7)},
		{NewList(1, 2, 3, 4, 5, 6, 7, 8), NewList(2, 1, 4, 3, 6, 5, 8, 7)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := swapPairs(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
