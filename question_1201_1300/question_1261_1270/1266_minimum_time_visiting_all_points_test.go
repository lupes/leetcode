package question_1261_1270

import (
	"testing"
)

func Test_minTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.points); got != tt.want {
				t.Errorf("minTimeToVisitAllPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
