package question_61_70

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{"test#0", 2, 1, 1},
		{"test#1", 1, 2, 1},
		{"test#2", 3, 2, 3},
		{"test#3", 7, 3, 28},
		{"test#4", 51, 9, 1916797311},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
