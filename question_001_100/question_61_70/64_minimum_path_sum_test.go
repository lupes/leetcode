package question_61_70

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"test#0", [][]int{{1}}, 1},
		{"test#1", [][]int{{1, 1}}, 2},
		{"test#2", [][]int{{1}, {1}}, 2},
		{"test#3", [][]int{{1, 2}, {2, 1}}, 4},
		{"test#4", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"test#5", [][]int{{1, 2, 5}, {3, 2, 1}}, 6},
		{"test#6", [][]int{{1, 2, 5}, {3, 0, 0}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
