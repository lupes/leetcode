package question_0011_0020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{NewList(1, 2, 3, 4, 5), 1, NewList(1, 2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 2, NewList(2, 1, 4, 3, 5)},
		{NewList(1, 2, 3, 4, 5), 3, NewList(3, 2, 1, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 4, NewList(4, 3, 2, 1, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(5, 4, 3, 2, 1)},
		{NewList(1), 2, NewList(1)},
		{nil, 6, nil},
		{NewList(1, 2), 2, NewList(2, 1)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := reverseKGroup(tt.head, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
