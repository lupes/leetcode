package question_981_990

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_smallestFromLeaf(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want string
	}{
		{NewNode(0, NewNode(1, NewNode(3), NewNode(4)), NewNode(2, NewNode(3), NewNode(4))), "dba"},
		{NewNode(25, NewNode(1, NewNode(1), NewNode(3)), NewNode(3, NewNode(0), NewNode(2))), "adz"},
		{NewNode(2, NewNode(2, nil, NewNode(1, NewNode(0))), NewNode(1, NewNode(0))), "abc"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := smallestFromLeaf(tt.root); got != tt.want {
				t.Errorf("smallestFromLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
