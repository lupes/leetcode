package question_1661_1670

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_mergeInBetween(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
		want  *ListNode
	}{
		{NewList(0, 1, 2, 3, 4, 5), 3, 4, NewList(1000000, 1000001, 1000002), NewList(0, 1, 2, 1000000, 1000001, 1000002, 5)},
		{NewList(0, 1, 2, 3, 4, 5, 6), 2, 5, NewList(1000000, 1000001, 1000002, 1000003, 1000004), NewList(0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mergeInBetween(tt.list1, tt.a, tt.b, tt.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
