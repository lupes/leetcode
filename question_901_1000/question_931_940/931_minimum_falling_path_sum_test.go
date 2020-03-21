package question_931_940

import (
	"testing"
)

func Test_minFallingPathSum(t *testing.T) {
	tests := []struct {
		A    [][]int
		want int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 12},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minFallingPathSum(tt.A); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
