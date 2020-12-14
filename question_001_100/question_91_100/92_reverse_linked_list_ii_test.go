package question_91_100

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_reverseBetween(t *testing.T) {
	head1 := &ListNode{}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	want2 := &ListNode{Val: 2, Next: &ListNode{Val: 1}}

	head3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want3 := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	head4 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want4 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4}}}}

	head5 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want5 := &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}

	head6 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want6 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}}

	head7 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want7 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}}

	head8 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	want8 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	tests := []struct {
		name string
		head *ListNode
		m    int
		n    int
		want *ListNode
	}{
		{"test", nil, 0, 0, nil},
		{"test", head1, 1, 1, head1},
		{"test", head2, 1, 2, want2},
		{"test", head3, 1, 2, want3},
		{"test", head4, 1, 3, want4},
		{"test", head5, 1, 4, want5},
		{"test", head6, 2, 3, want6},
		{"test", head7, 2, 4, want7},
		{"test", head8, 4, 4, want8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.head, tt.m, tt.n)
			fmt.Print("got:  ")
			for next := got; next != nil; next = next.Next {
				fmt.Printf("%d ", next.Val)
			}
			fmt.Println()
			fmt.Print("want: ")
			for next := tt.want; next != nil; next = next.Next {
				fmt.Printf("%d ", next.Val)
			}
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got:%v want:%v\n", got, tt.want)
			}
		})
	}
}
