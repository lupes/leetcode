package question_1001_1010

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_bstFromPreorder(t *testing.T) {
	tests := []struct {
		preorder []int
		want     *TreeNode
	}{
		{[]int{8, 5, 1, 7, 10, 12}, NewNode(8, NewNode(5, NewNode(1), NewNode(7)), NewNode(10, nil, NewNode(12)))},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := bstFromPreorder(tt.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
