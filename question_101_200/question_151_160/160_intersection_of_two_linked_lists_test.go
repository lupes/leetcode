package question_151_160

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	end1 := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA1 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: end1}}
	headB1 := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: end1}}}

	end2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	headA2 := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: end2}}}
	headB2 := &ListNode{Val: 3, Next: end2}

	headA3 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	headB3 := &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	tests := []struct {
		name  string
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		{"test", headA1, headB1, end1},
		{"test", headA2, headB2, end2},
		{"test", headA3, headB3, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.headA, tt.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
