package question_0011_0020

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

var lists1 = []*ListNode{
	{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
	{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	{Val: 2, Next: &ListNode{Val: 6}},
}

var want1 = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}}}}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test#1", args{lists1}, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
