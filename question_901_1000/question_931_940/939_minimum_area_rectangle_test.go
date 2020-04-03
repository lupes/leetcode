package question_931_940

import (
	"testing"
)

func Test_minAreaRect(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}, 4},
		{[][]int{{1, 1}, {1, 3}}, 0},
		{[][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minAreaRect(tt.points); got != tt.want {
				t.Errorf("minAreaRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
