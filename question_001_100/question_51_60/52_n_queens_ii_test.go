package question_51_60

import "testing"

func Test_totalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test#0", 1, 1},
		{"test#1", 2, 0},
		{"test#2", 3, 0},
		{"test#3", 4, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
