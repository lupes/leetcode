package question_81_90

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	want1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}}
	tests := []struct {
		name string
		head *ListNode
		x    int
		want *ListNode
	}{
		{"test#0", nil, 1, nil},
		{"test#1", &ListNode{Val: 1}, 0, &ListNode{Val: 1}},
		{"test#2", head1, 3, want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.head, tt.x)
			for n := got; n != nil; n = n.Next {
				fmt.Printf("%d ", n.Val)
			}
			fmt.Println()
			for n := tt.want; n != nil; n = n.Next {
				fmt.Printf("%d ", n.Val)
			}
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
