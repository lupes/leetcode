package question_221_230

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_invertTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{nil, nil},
		{NewNode(4), NewNode(4)},
		{NewNode(4, NewNode(2)), NewNode(4, nil, NewNode(2))},
		{NewNode(4, NewNode(2, NewNode(1), NewNode(3)), NewNode(7, NewNode(6), NewNode(9))),
			NewNode(4, NewNode(7, NewNode(9), NewNode(6)), NewNode(2, NewNode(3), NewNode(1)))},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := invertTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
