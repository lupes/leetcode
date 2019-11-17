package question_71_80

import "testing"

func Test_searchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"test#0", [][]int{}, 0, false},
		{"test#0", [][]int{{}}, 0, false},
		{"test#1", [][]int{{1}}, 0, false},
		{"test#2", [][]int{{1, 2}}, 1, true},
		{"test#3", [][]int{{1, 2, 3}}, 4, false},
		{"test#4", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3, true},
		{"test#5", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
