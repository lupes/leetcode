package question_01_10

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{NewList(2, 4, 3), NewList(5, 6, 4), NewList(7, 0, 8)},
		{NewList(2), NewList(5), NewList(7)},
		{NewList(5), NewList(5), NewList(0, 1)},
		{NewList(5), NewList(5, 1), NewList(0, 2)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
