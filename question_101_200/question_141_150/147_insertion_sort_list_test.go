package question_141_150

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	head1 := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
	want1 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	head2 := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	want2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	head3 := &ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0}}}}}
	want3 := &ListNode{Val: -1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		//{"test", nil, nil},
		//{"test", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"test", head1, want1},
		{"test", head2, want2},
		{"test", head3, want3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertionSortList(tt.head)
			for n := got; n != nil; n = n.Next {
				fmt.Printf("%d ", n.Val)
			}
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
