package question_1881_1890

import (
	"testing"
)

func Test_minFlips(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"111000", 2},
		{"010", 0},
		{"1110", 1},
		{"11101", 1},
		{"111011", 2},
		{"1110101", 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minFlips(tt.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
