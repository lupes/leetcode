package question_591_600

import "testing"

func Test_maxCount(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		ops  [][]int
		want int
	}{
		{3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{3, 3, [][]int{{2, 1}, {3, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxCount(tt.m, tt.n, tt.ops); got != tt.want {
				t.Errorf("maxCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
