package question_1611_1620

import (
	"testing"
)

func Test_maximalNetworkRank(t *testing.T) {
	tests := []struct {
		n     int
		roads [][]int
		want  int
	}{
		{4, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}}, 4},
		{5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}}, 5},
		{8, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximalNetworkRank(tt.n, tt.roads); got != tt.want {
				t.Errorf("maximalNetworkRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
