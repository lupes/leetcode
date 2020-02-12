package question_811_820

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_pruneTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want *TreeNode
	}{
		{NewNode(1, nil, NewNode(0, NewNode(0), NewNode(1))),
			NewNode(1, nil, NewNode(0, nil, NewNode(1)))},
		{NewNode(1, NewNode(0, NewNode(0), NewNode(0)), NewNode(1, NewNode(0), NewNode(1))),
			NewNode(1, nil, NewNode(1, nil, NewNode(1)))},
		{NewNode(1, NewNode(0, NewNode(0), NewNode(1)), NewNode(1, NewNode(0), NewNode(1))),
			NewNode(1, NewNode(0, nil, NewNode(1)), NewNode(1, nil, NewNode(1)))},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pruneTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
