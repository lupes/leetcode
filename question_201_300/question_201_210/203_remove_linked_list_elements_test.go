package question_201_210

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		val  int
		want *ListNode
	}{
		{"test", nil, 1, nil},
		{"test", &ListNode{Val: 1}, 1, nil},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1, &ListNode{Val: 2}},
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 1}}, 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.head, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
