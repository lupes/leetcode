package question_81_90

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"test#0", nil, nil},
		{"test#1", &ListNode{Val: 1}, &ListNode{Val: 1}},
		{"test#2", &ListNode{Val: 1, Next: &ListNode{Val: 1}}, &ListNode{Val: 1}},
		{"test#3", &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{"test#4", &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}, &ListNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
