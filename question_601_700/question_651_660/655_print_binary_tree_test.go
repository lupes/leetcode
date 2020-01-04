package question_651_660

import (
	"reflect"
	"testing"

	. "github.com/lupes/leetcode/common"
)

var (
	root1 = NewNode(1)
	root2 = NewNode(1, NewNode(2))
	root3 = NewNode(1, NewNode(2, nil, NewNode(4)), NewNode(3))
)

func Test_depth(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{root1, 1},
		{root2, 2},
		{root3, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := depth(tt.root); got != tt.want {
				t.Errorf("depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]string
	}{
		{root1, [][]string{{"1"}}},
		{root2, [][]string{{"", "1", ""}, {"2", "", ""}}},
		{root3, [][]string{{"", "", "", "1", "", "", ""}, {"", "2", "", "", "", "3", ""}, {"", "", "4", "", "", "", ""}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := printTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printTree() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
