package question_1301_1310

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_getAllElements(t *testing.T) {
	tests := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  []int
	}{
		{NewNodeV2(2, 1, 4), NewNodeV2(1, 0, 3), []int{0, 1, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getAllElements(tt.root1, tt.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
