package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	var head0 *ListNode
	var head1 = &ListNode{Val: 1, Next: nil}
	var head2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	var head3 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	var head4 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	var head5 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	var want0 *ListNode
	var want1 = &ListNode{Val: 1, Next: nil}
	var want2 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}
	var want3 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}
	var want4 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}}
	var want5 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}}

	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"test#0", head0, want0},
		{"test#1", head1, want1},
		{"test#2", head2, want2},
		{"test#3", head3, want3},
		{"test#4", head4, want4},
		{"test#5", head5, want5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.head)
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
				t.Errorf("swapPairs() = %v, want %v", arr1, arr2)
			}
		})
	}
}
