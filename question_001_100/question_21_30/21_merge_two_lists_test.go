package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	var (
		t0l1 *ListNode
		t0l2 *ListNode
		t0w  *ListNode

		t1l1 = &ListNode{Val: 1}
		t1l2 *ListNode
		t1w  = &ListNode{Val: 1}

		t2l1 *ListNode
		t2l2 = &ListNode{Val: 1}
		t2w  = &ListNode{Val: 1}

		t3l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
		t3l2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
		t3w  = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}
	)
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{"test#0", t0l1, t0l2, t0w},
		{"test#1", t1l1, t1l2, t1w},
		{"test#2", t2l1, t2l2, t2w},
		{"test#3", t3l1, t3l2, t3w},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.l1, tt.l2)
			want := tt.want
			var arr1 []int
			var arr2 []int
			for got != nil {
				arr1 = append(arr1, got.Val)
				got = got.Next
			}

			for want != nil {
				arr2 = append(arr2, want.Val)
				want = want.Next
			}
			if !reflect.DeepEqual(arr1, arr2) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
