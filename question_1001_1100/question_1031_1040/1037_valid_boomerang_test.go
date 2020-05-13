package question_1031_1040

import (
	"testing"
)

func Test_isBoomerang(t *testing.T) {
	tests := []struct {
		points [][]int
		want   bool
	}{
		{[][]int{{1, 1}, {2, 3}, {3, 2}}, true},
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, false},
		{[][]int{{0, 0}, {1, 0}, {2, 2}}, true},
		{[][]int{{0, 0}, {0, 1}, {2, 2}}, true},
		{[][]int{{0, 1}, {1, 2}, {2, 2}}, true},
		{[][]int{{1, 1}, {3, 3}, {2, 2}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isBoomerang(tt.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}
