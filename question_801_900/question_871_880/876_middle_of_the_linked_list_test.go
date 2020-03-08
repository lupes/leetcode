package question_871_880

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{NewList(1), NewList(1)},
		{NewList(1, 2), NewList(2)},
		{NewList(1, 2, 3), NewList(2, 3)},
		{NewList(1, 2, 3, 4), NewList(3, 4)},
		{NewList(1, 2, 3, 4, 5), NewList(3, 4, 5)},
		{NewList(1, 2, 3, 4, 5, 6), NewList(4, 5, 6)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := middleNode(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
