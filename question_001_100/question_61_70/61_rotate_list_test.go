package question_61_70

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want *ListNode
	}{
		{"test#0", nil, 1, nil},
		{"test#1", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2, &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}}},
		{"test#2", &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, 4, &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.head, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
