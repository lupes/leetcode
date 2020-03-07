package question_871_880

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_leafSimilar(t *testing.T) {
	tests := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{NewNode(1), NewNode(1), true},
		{NewNode(1, NewNode(2), NewNode(3)), NewNode(1), false},
		{NewNode(1, NewNode(4, NewNode(2)), NewNode(3)), NewNode(1, NewNode(2), NewNode(3)), true},
		{NewNode(3, NewNode(5, NewNode(6), NewNode(2, NewNode(7), NewNode(4))), NewNode(1, NewNode(9), NewNode(8))), NewNode(1, NewNode(2), NewNode(3)), false},
		{NewNode(3, NewNode(5, NewNode(6), NewNode(2, NewNode(7), NewNode(4))), NewNode(1, NewNode(9))), NewNode(1, NewNode(2, NewNode(6), NewNode(7)), NewNode(3, NewNode(4), NewNode(9))), true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := leafSimilar(tt.root1, tt.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
