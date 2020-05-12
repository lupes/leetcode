package question_551_560

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		root *Node
		want int
	}{
		{&Node{Val: 1, Children: []*Node{{Val: 3, Children: []*Node{{Val: 5}}}, {Val: 2}}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
