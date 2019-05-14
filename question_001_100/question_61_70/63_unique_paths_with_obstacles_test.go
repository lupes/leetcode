package question_61_70

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{"test#0", [][]int{{1}}, 0},
		{"test#1", [][]int{{0, 0}}, 1},
		{"test#2", [][]int{{1, 0}}, 0},
		{"test#3", [][]int{{0}, {0}}, 1},
		{"test#4", [][]int{{0, 1}, {0, 0}}, 1},
		{"test#5", [][]int{{0, 0}, {0, 0}}, 2},
		{"test#6", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{"test#7", [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
