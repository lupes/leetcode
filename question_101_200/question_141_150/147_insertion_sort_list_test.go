package question_141_150

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_insertionSortList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		//{"test", nil, nil},
		//{"test", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{NewList(2, 1), NewList(1, 2)},
		{NewList(4, 2, 1, 3), NewList(1, 2, 3, 4)},
		{NewList(-1, 5, 3, 4, 0), NewList(-1, 0, 3, 4, 5)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := insertionSortList(tt.head)
			fmt.Printf("%+v\n", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
