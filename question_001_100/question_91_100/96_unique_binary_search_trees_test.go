package question_91_100

import "testing"

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 5},
		{17, 129644790},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
