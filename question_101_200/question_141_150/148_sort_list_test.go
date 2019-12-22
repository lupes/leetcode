package question_141_150

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_sortList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{nil, nil},
		{NewList(4, 2, 1, 3), NewList(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
