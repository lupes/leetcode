package question_91_100

import (
	"fmt"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
	}{
		{NewNodeV2(1, 3, Null, Null, 2)},
		{NewNodeV2(3, 1, 4, Null, Null, 2)},
		{NewNodeV2(146, 71, -13, 55, Null, 231, 399, 321, Null, Null, Null, Null, Null, -33)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			fmt.Printf("%v\n", inorderTraversal(tt.root))
			recoverTree(tt.root)
			fmt.Printf("%v\n", inorderTraversal(tt.root))
		})
	}
}
