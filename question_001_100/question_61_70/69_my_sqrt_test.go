package question_61_70

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test#0", 0, 0},
		{"test#1", 1, 1},
		{"test#2", 2, 1},
		{"test#3", 3, 1},
		{"test#4", 4, 2},
		{"test#5", 5, 2},
		{"test#6", 6, 2},
		{"test#7", 7, 2},
		{"test#8", 8, 2},
		{"test#9", 9, 3},
		{"test#10", 254, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
