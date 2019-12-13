package question_221_230

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_countNodes(t *testing.T) {
	root1 := NewNode(1)
	root2 := NewNode(1, NewNode(21), NewNode(22))
	root3 := NewNode(1, NewNode(21, NewNode(31), NewNode(32)), NewNode(22, NewNode(33), NewNode(34)))
	root4 := NewNode(1,
		NewNode(21, NewNode(31, NewNode(41), NewNode(42)), NewNode(32, NewNode(43), NewNode(44))),
		NewNode(22, NewNode(33, NewNode(45), NewNode(46)), NewNode(34, NewNode(47), NewNode(48))))
	root5 := NewNode(1,
		NewNode(21,
			NewNode(31, NewNode(41, NewNode(501), NewNode(502)), NewNode(42, NewNode(503), NewNode(504))),
			NewNode(32, NewNode(43, NewNode(505), NewNode(506)), NewNode(44, NewNode(507), NewNode(508)))),
		NewNode(22,
			NewNode(33, NewNode(45, NewNode(509), NewNode(510)), NewNode(46, NewNode(511), NewNode(512))),
			NewNode(34, NewNode(47, NewNode(513), NewNode(514)), NewNode(48, NewNode(515), NewNode(516)))))
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"test", nil, 0},
		{"test", root1, 1},
		{"test", root2, 3},
		{"test", root3, 7},
		{"test", root4, 15},
		{"test", root5, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
