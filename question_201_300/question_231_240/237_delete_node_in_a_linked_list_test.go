package question_231_240

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		name string
		node *ListNode
		want *ListNode
	}{
		{"test", &ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.node)
			if !reflect.DeepEqual(tt.node, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.node, tt.want)
			}
		})
	}
}
