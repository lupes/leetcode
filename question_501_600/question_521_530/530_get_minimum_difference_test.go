package question_521_530

import "testing"

func Test_getMinimumDifference(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}}, 1},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getMinimumDifference(tt.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMinimumDifferenceDfs(t *testing.T) {
	tests := []struct {
		root *TreeNode
		min  int
		max  int
		res  int
	}{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}}, 0, 1, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			min, max, res := getMinimumDifferenceDfs(tt.root)
			if min != tt.min {
				t.Errorf("getMinimumDifferenceDfs() got min = %v, want min = %v", min, tt.min)
			}
			if max != tt.max {
				t.Errorf("getMinimumDifferenceDfs() got max = %v, want max = %v", max, tt.max)
			}
			if res != tt.res {
				t.Errorf("getMinimumDifferenceDfs() got res = %v, want res = %v", res, tt.res)
			}
		})
	}
}
