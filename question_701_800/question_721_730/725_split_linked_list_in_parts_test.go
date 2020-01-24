package question_721_730

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_splitListToParts(t *testing.T) {
	tests := []struct {
		root *ListNode
		k    int
		want []*ListNode
	}{
		{NewList(1, 2, 3), 5,
			[]*ListNode{NewList(1), NewList(2), NewList(3), nil, nil}},
		{NewList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 3,
			[]*ListNode{NewList(1, 2, 3, 4), NewList(5, 6, 7), NewList(8, 9, 10)}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := splitListToParts(tt.root, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
