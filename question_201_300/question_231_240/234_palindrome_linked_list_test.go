package question_231_240

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"test", nil, true},
		{"test", &ListNode{Val: 1}, true},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, false},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}, true},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
