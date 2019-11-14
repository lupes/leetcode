package question_11_20

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{nil, 1, nil},
		{NewList(1), 1, nil},
		{NewList(1, 2, 3, 4, 5), 2, NewList(1, 2, 3, 5)},
		{NewList(1, 2, 3, 4, 5), 5, NewList(2, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
