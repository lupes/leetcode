package question_501_510

import (
	"reflect"
	"sort"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func Test_findFrequentTreeSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"test", &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -3}}, []int{-3, 2, 4}},
		{"test", &TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -5}}, []int{2}},
		{"test", &TreeNode{Val: 5, Left: &TreeNode{Val: 2}}, []int{2, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFrequentTreeSum(tt.root)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
