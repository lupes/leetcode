package question_321_330

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_oddEvenList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{NewList(1, 2, 3, 4, 5), NewList(1, 3, 5, 2, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := oddEvenList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
