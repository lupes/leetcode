package question_0011_0020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{
			[]*ListNode{NewList(1, 4, 5), NewList(1, 3, 4), NewList(2, 6)},
			NewList(1, 1, 2, 3, 4, 4, 5, 6),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mergeKLists(tt.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
