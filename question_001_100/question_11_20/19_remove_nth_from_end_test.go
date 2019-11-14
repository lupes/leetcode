package question_11_20

import (
	"reflect"
	"testing"
)

var head1 *ListNode
var head2 = &ListNode{Val: 1, Next: nil}
var head3 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
var res3 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}}

var head4 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
var res4 = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		{head1, 1, nil},
		{head2, 2, nil},
		{head3, 2, res3},
		{head4, 5, res4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
