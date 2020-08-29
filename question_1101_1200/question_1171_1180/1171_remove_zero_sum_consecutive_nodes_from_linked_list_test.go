package question_1171_1180

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_removeZeroSumSublists(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{NewList(1, 2, 3, 1), NewList(1, 2, 3, 1)},
		{NewList(1, 2, -3, 3, 1), NewList(3, 1)},
		{NewList(1, 2, 3, -3, 4), NewList(1, 2, 4)},
		{NewList(1, 2, 3, -3, -2), NewList(1)},
		{NewList(1, 2, 3, -3, -2, -1), nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeZeroSumSublists(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeZeroSumSublists() = %v, want %v", got, tt.want)
			}
		})
	}
}
