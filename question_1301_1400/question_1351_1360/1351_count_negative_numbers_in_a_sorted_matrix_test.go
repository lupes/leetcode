package question_1331_1340

import (
	"testing"
)

func Test_countNegatives(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countNegatives(tt.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
