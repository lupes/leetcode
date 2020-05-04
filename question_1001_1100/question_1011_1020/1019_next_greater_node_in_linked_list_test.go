package question_1011_1020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_nextLargerNodes(t *testing.T) {
	tests := []struct {
		head *ListNode
		want []int
	}{
		{NewList(2, 1, 5), []int{5, 5, 0}},
		{NewList(2, 7, 4, 3, 5), []int{7, 0, 5, 5, 0}},
		{NewList(1, 7, 5, 1, 9, 2, 5, 1), []int{7, 9, 9, 9, 0, 5, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := nextLargerNodes(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextLargerNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
