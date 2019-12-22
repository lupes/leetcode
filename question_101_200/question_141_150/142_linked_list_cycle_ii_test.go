package question_141_150

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_detectCycle(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}

	n1.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n3
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"test", nil, nil},
		{"test", &ListNode{Val: 1}, nil},
		{"test", n1, n3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
