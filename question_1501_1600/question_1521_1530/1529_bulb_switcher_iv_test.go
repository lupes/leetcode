package question_1521_1530

import (
	"testing"
)

func Test_minFlips(t *testing.T) {
	tests := []struct {
		target string
		want   int
	}{
		{"1", 1},
		{"0", 0},
		{"10111", 3},
		{"101", 3},
		{"00000", 0},
		{"001011101", 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minFlips(tt.target); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
