package question_201_210

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"test", nil, nil},
		{"test", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
