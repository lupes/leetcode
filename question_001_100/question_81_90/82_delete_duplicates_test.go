package question_81_90

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"test#0", nil, nil},
		{"test#1", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"test#2", &ListNode{Val: 1, Next: &ListNode{Val: 1}}, nil},
		{"test#3", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{"test#4", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}},
		{"test#5", &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}, &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		{"test#6", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}, &ListNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates2(tt.head)
			for n := tt.want; n != nil; n = n.Next {
				fmt.Printf("%d ", n.Val)
			}
			fmt.Println()
			for n := got; n != nil; n = n.Next {
				fmt.Printf("%d ", n.Val)
			}
			fmt.Println()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
