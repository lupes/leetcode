package question_801_810

import (
	"testing"
)

func Test_maxIncreaseKeepingSkyline(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}, 35},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxIncreaseKeepingSkyline(tt.grid); got != tt.want {
				t.Errorf("maxIncreaseKeepingSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
