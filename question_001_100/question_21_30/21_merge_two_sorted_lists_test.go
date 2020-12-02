package question_0011_0020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{nil, nil, nil},
		{NewList(1), nil, NewList(1)},
		{nil, NewList(1), NewList(1)},
		{NewList(1, 2, 4), NewList(1, 3, 4), NewList(1, 1, 2, 3, 4, 4)},
		{NewList(1, 3), NewList(2, 4), NewList(1, 2, 3, 4)},
		{NewList(1, 3, 3), NewList(2, 4), NewList(1, 2, 3, 3, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := mergeTwoLists(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
