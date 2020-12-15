package question_91_100

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{nil, 0, 0, nil},
		{NewList(), 1, 1, NewList()},
		{NewList(1, 2), 1, 2, NewList(2, 1)},
		{NewList(1, 2, 3, 4), 1, 2, NewList(2, 1, 3, 4)},
		{NewList(1, 2, 3, 4), 1, 3, NewList(3, 2, 1, 4)},
		{NewList(1, 2, 3, 4), 1, 4, NewList(4, 3, 2, 1)},
		{NewList(1, 2, 3, 4), 2, 3, NewList(1, 3, 2, 4)},
		{NewList(1, 2, 3, 4), 2, 4, NewList(1, 4, 3, 2)},
		{NewList(1, 2, 3, 4), 4, 4, NewList(1, 2, 3, 4)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := reverseBetween(tt.head, tt.m, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got:%v want:%v\n", got, tt.want)
			}
		})
	}
}
