package question_1031_1040

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_bstToGst(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{NewNode(2, NewNode(1), NewNode(3)), NewNode(5, NewNode(6), NewNode(3))},
		{
			NewNode(4, NewNode(1, NewNode(0), NewNode(2, nil, NewNode(3))),
				NewNode(6, NewNode(5), NewNode(7, nil, NewNode(8)))),
			NewNode(30, NewNode(36, NewNode(36), NewNode(35, nil, NewNode(33))),
				NewNode(21, NewNode(26), NewNode(15, nil, NewNode(8)))),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := bstToGst(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstToGst() = \n     %v, \nwant %v", got, tt.want)
			}
		})
	}
}
