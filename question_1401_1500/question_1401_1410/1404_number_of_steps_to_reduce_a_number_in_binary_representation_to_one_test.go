package question_1401_1410

import (
	"testing"
)

func Test_numSteps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"1101", 6},
		{"10", 1},
		{"1", 0},
		{"1111011110000011100000110001011011110010111001010111110001", 85},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSteps(tt.s); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
