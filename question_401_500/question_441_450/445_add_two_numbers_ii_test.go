package question_441_450

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{"test",
			&ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 7}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				for n := got; n != nil; n = n.Next {
					fmt.Printf("%d ", n.Val)
				}
				fmt.Println()
				for n := tt.want; n != nil; n = n.Next {
					fmt.Printf("%d ", n.Val)
				}
				fmt.Println()
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
