package question_611_620

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_mergeTrees(t *testing.T) {
	tests := []struct {
		t1   *TreeNode
		t2   *TreeNode
		want *TreeNode
	}{
		{
			nil, nil, nil,
		},
		{
			NewNode(1, NewNode(3, NewNode(5)), NewNode(2)),
			nil,
			NewNode(1, NewNode(3, NewNode(5)), NewNode(2)),
		},
		{
			nil,
			NewNode(1, NewNode(3, NewNode(5)), NewNode(2)),
			NewNode(1, NewNode(3, NewNode(5)), NewNode(2)),
		},
		{
			NewNode(1, NewNode(3, NewNode(5)), NewNode(2)),
			NewNode(2, NewNode(1, nil, NewNode(4)), NewNode(3, nil, NewNode(7))),
			NewNode(3, NewNode(4, NewNode(5), NewNode(4)), NewNode(5, nil, NewNode(7))),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mergeTrees(tt.t1, tt.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
