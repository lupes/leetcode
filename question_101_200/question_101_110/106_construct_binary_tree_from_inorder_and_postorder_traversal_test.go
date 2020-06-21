package question_101_110

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_buildTree106(t *testing.T) {
	tests := []struct {
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, NewNodeV2(3, 9, 20, math.MaxInt32, math.MaxInt32, 15, 7)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := buildTree106(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree106() = %v, want %v", got, tt.want)
			}
		})
	}
}
