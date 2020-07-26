package question_621_630

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		root *TreeNode
		v    int
		d    int
		want *TreeNode
	}{
		{NewNodeV2(4, 2, 6, 3, 1, 5), 1, 2,
			NewNodeV2(4, 1, 1, 2, math.MaxInt32, math.MaxInt32, 6, 3, 1, 5),
		},
		{
			NewNodeV2(4, 2, math.MaxInt32, 3, 1), 1, 3,
			NewNodeV2(4, 2, math.MaxInt32, 1, 1, 3, math.MaxInt32, math.MaxInt32, 1),
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := addOneRow(tt.root, tt.v, tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
