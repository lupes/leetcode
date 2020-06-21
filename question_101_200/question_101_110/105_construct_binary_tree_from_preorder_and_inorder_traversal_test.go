package question_101_110

import (
	"math"
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, NewNodeV2(3, 9, 20, math.MaxInt32, math.MaxInt32, 15, 7)},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := buildTree(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
