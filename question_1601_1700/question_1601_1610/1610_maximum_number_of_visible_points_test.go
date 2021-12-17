package question_1601_1610

import (
	"testing"
)

func Test_visiblePoints(t *testing.T) {
	tests := []struct {
		points   [][]int
		angle    int
		location []int
		want     int
	}{
		{[][]int{{2, 1}, {2, 2}, {3, 3}}, 90, []int{1, 1}, 3},
		{[][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}}, 90, []int{1, 1}, 4},
		{[][]int{{1, 0}, {2, 1}}, 13, []int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := visiblePoints(tt.points, tt.angle, tt.location); got != tt.want {
				t.Errorf("visiblePoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
