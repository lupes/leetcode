package question_601_610

import "testing"

func Test_tree2str(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		t    *TreeNode
		want string
	}{
		{nil, ""},
		{&TreeNode{Val: 0}, "0"},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, "1(2(4))(3)"},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, "1(2()(4))(3)"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := tree2str(tt.t); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
