package question_671_680

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_findSecondMinimumValue(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{NewNode(2, NewNode(2), NewNode(5, NewNode(5), NewNode(7))), 5},
		{NewNode(2, NewNode(2), NewNode(5, NewNode(5), NewNode(5))), 5},
		{NewNode(2, NewNode(2), NewNode(2)), -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findSecondMinimumValue(tt.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
