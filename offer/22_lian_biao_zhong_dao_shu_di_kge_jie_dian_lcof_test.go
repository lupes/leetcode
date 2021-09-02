package offer

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_getKthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		want *ListNode
	}{
		{NewList(1, 2, 3, 4, 5), 1, NewList(5)},
		{NewList(1, 2, 3, 4, 5), 2, NewList(4, 5)},
		{NewList(1, 2, 3, 4, 5), 3, NewList(3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 4, NewList(2, 3, 4, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(1, 2, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getKthFromEnd(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
