package question_861_870

import (
	"testing"
)

func Test_matrixScore(t *testing.T) {
	tests := []struct {
		A    [][]int
		want int
	}{
		{[][]int{{1}}, 1},
		{[][]int{{0}}, 1},
		{[][]int{{0, 1}}, 3},
		{[][]int{{0, 1}, {1, 0}}, 6},
		{[][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, 39},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := matrixScore(tt.A); got != tt.want {
				t.Errorf("matrixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
