package question_1571_1580

import (
	"testing"
)

func Test_numWays(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"10101", 4},
		{"0000", 3},
		{"1001", 0},
		{"100100010100110", 12},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numWays(tt.s); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
