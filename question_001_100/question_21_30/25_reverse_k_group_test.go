package question_0011_0020

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {

	var head1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var head2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var head3 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var head4 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var head5 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var head6 = &ListNode{Val: 1, Next: nil}
	var head7 *ListNode
	var head8 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}

	var want2 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}}
	var want3 = &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	var want4 = &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}}}}
	var want5 = &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}
	var want7 *ListNode
	var want8 = &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}

	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test#1", args{head: head1, k: 1}, head1},
		{"test#2", args{head: head2, k: 2}, want2},
		{"test#3", args{head: head3, k: 3}, want3},
		{"test#4", args{head: head4, k: 4}, want4},
		{"test#5", args{head: head5, k: 5}, want5},
		{"test#6", args{head: head6, k: 2}, head6},
		{"test#7", args{head: head7, k: 6}, want7},
		{"test#8", args{head: head8, k: 2}, want8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			want := tt.want
			var arr1 []int
			var arr2 []int
			for got != nil && want != nil {
				arr1 = append(arr1, got.Val)
				arr2 = append(arr2, want.Val)
				got = got.Next
				want = want.Next
			}

			if !reflect.DeepEqual(arr1, arr2) {
				t.Errorf("reverseKGroup() = %v, want %v", arr1, arr2)
			}
		})
	}
}
