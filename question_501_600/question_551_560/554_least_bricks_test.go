package question_551_560

import "testing"

func Test_leastBricks(t *testing.T) {
	tests := []struct {
		wall [][]int
		want int
	}{
		{[][]int{{1}, {1}, {1}}, 3},
		{[][]int{{2}, {2}, {2}}, 3},
		{[][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := leastBricks(tt.wall); got != tt.want {
				t.Errorf("leastBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
