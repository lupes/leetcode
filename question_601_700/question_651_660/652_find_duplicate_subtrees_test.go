package question_651_660

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

//     0
//   0   0
//  0 n  n 0
// n n    0 n

func Test_findDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []*TreeNode
	}{
		{
			NewNode(1, NewNode(2, NewNode(4)), NewNode(3, NewNode(2, NewNode(4)), NewNode(4))),
			[]*TreeNode{NewNode(4), NewNode(2, NewNode(4))},
		},
		{
			NewNode(2, NewNode(2, NewNode(3)), NewNode(2, NewNode(3))),
			[]*TreeNode{NewNode(3), NewNode(2, NewNode(3))},
		},
		{
			NewNode(0, NewNode(0, NewNode(0)), NewNode(0, nil, NewNode(0, NewNode(0), NewNode(0)))),
			[]*TreeNode{NewNode(0)},
		},
		{
			NewNode(0, NewNode(0, NewNode(0)), NewNode(0, nil, NewNode(0, NewNode(0), NewNode(0)))),
			[]*TreeNode{NewNode(0)},
		},
		{
			NewNode(0, NewNode(0, NewNode(0, NewNode(0), NewNode(0))), NewNode(0, nil, NewNode(0, NewNode(0), NewNode(0)))),
			[]*TreeNode{NewNode(0), NewNode(0, NewNode(0), NewNode(0))},
		},
		{
			NewNode(0, NewNode(0, NewNode(0)), NewNode(0, nil, NewNode(0, NewNode(0)))),
			[]*TreeNode{NewNode(0), NewNode(0, NewNode(0))},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
