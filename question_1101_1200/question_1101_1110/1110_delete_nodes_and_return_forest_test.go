package question_1101_1110

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_delNodes(t *testing.T) {
	tests := []struct {
		root      *TreeNode
		to_delete []int
		want      []*TreeNode
	}{
		{NewNodeV2(1, 2, 3, 4, 5, 6, 7), []int{3, 5}, []*TreeNode{
			NewNode(1, NewNode(2, NewNode(4))), NewNode(6), NewNode(7),
		}},
		{NewNode(1, NewNode(2), NewNode(3, nil, NewNode(4))), []int{1, 2}, []*TreeNode{
			NewNode(3, nil, NewNode(4)),
		}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := delNodes(tt.root, tt.to_delete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
