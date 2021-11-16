package question_391_400

import (
	"testing"
)

func Test_isRectangleCover(t *testing.T) {
	tests := []struct {
		rectangles [][]int
		want       bool
	}{
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}, true},
		{[][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}, false},
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {3, 2, 4, 4}}, false},
		{[][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {1, 3, 2, 4}, {2, 2, 4, 4}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isRectangleCover(tt.rectangles); got != tt.want {
				t.Errorf("isRectangleCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
