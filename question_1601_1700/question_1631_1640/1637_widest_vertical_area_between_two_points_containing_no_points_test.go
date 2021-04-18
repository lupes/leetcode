package question_1631_1640

import (
	"testing"
)

func Test_maxWidthOfVerticalArea(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}, 1},
		{[][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxWidthOfVerticalArea(tt.points); got != tt.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
