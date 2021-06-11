package question_441_450

import (
	"testing"
)

func Test_numberOfBoomerangs(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{[][]int{{0, 0}, {1, 0}, {2, 0}}, 2},
		{[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 4},
		{[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfBoomerangs(tt.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
