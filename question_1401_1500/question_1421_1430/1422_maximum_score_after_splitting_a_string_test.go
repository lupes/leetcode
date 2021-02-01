package question_1421_1430

import (
	"testing"
)

func Test_maxScore1422(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
		{"0000", 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxScore1422(tt.s); got != tt.want {
				t.Errorf("maxScore1423() = %v, want %v", got, tt.want)
			}
		})
	}
}
